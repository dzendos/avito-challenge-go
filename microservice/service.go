package microservice

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/dzendos/avito-challenge/api"
	"github.com/dzendos/avito-challenge/internal/config"
	"github.com/dzendos/avito-challenge/internal/types"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	cancelCode  = -1
	approveCode = 1
)

type bankAccountDB interface {
	AddRefill(ctx context.Context, userID int64, amount int64) error
	GetAmount(ctx context.Context, userID int64) (int64, error)
	GetBalanceHistory(ctx context.Context, userID int64) ([]types.BalanceHistoryUnit, error)
}

type operationsDB interface {
	AddOperation(ctx context.Context, userID, serviceID, orderID, amount int64) error
	ModifyOperation(ctx context.Context, userID, serviceID, orderID, amount int64, code int) error
}

type server struct {
	api.UnimplementedQueryListenerServer
	logger        *zap.Logger
	bankAccountDB bankAccountDB
	operationsDB  operationsDB
}

func (s server) PostCredit(ctx context.Context, in *api.CreditInfo) (*api.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"PostCredit",
	)
	defer span.Finish()

	err := s.bankAccountDB.AddRefill(ctx, in.GetUserId(), in.GetAmount())
	if err != nil {
		s.logger.Error("cannot AddRefill", zap.Error(err))
		return &api.Empty{}, errors.Wrap(err, "cannot AddRefill")
	}

	return &api.Empty{}, nil
}

func (s server) PostGetBalance(ctx context.Context, in *api.User) (*api.Balance, error) {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"PostGetBalance",
	)
	defer span.Finish()

	amount, err := s.bankAccountDB.GetAmount(ctx, in.GetUserId())

	if err != nil {
		s.logger.Error("cannot GetAmount", zap.Error(err))
		return &api.Balance{Amount: 0}, errors.Wrap(err, "cannot GetAmount")
	}

	return &api.Balance{Amount: amount}, nil
}

func (s server) PostReserve(ctx context.Context, in *api.CashFlow) (*api.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"PostReserve",
	)
	defer span.Finish()

	amount, err := s.bankAccountDB.GetAmount(ctx, in.GetUserId())

	if err != nil {
		s.logger.Error("cannot GetAmount", zap.Error(err))
		return &api.Empty{}, errors.Wrap(err, "cannot GetAmount")
	}

	// In this case we want to reserve more money than we have
	if in.GetAmount() > amount {
		return &api.Empty{}, errors.New("Not enough money on user's credit")
	}

	// We assume that amount is positive,
	// then we add this amount to the existing one with the negative sign
	// we register negative movement in the history
	err = s.bankAccountDB.AddRefill(ctx, in.GetUserId(), -in.GetAmount())
	if err != nil {
		s.logger.Error("cannot AddRefill", zap.Error(err))
		return &api.Empty{}, errors.Wrap(err, "cannot AddRefill")
	}

	err = s.operationsDB.AddOperation(
		ctx,
		in.GetUserId(),
		in.GetServiceId(),
		in.GetOrderId(),
		in.GetAmount(),
	)
	if err != nil {
		s.logger.Error("cannot AddOperation", zap.Error(err))
		return &api.Empty{}, errors.Wrap(err, "cannot AddOperation")
	}

	return &api.Empty{}, nil
}

func (s server) PostCancelReserve(ctx context.Context, in *api.CashFlow) (*api.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"PostCancelReserve",
	)
	defer span.Finish()

	err := s.operationsDB.ModifyOperation(
		ctx,
		in.GetUserId(),
		in.GetServiceId(),
		in.GetOrderId(),
		in.GetAmount(),
		cancelCode,
	)

	if err != nil {
		s.logger.Error("cannot ModifyOperation", zap.Error(err))
		return &api.Empty{}, errors.Wrap(err, "cannot ModifyOperation")
	}

	err = s.bankAccountDB.AddRefill(ctx, in.GetUserId(), in.GetAmount())
	if err != nil {
		s.logger.Error("cannot AddRefill", zap.Error(err))
		return &api.Empty{}, errors.Wrap(err, "cannot AddRefill")
	}

	return &api.Empty{}, nil
}

func (s server) PostWriteOff(ctx context.Context, in *api.CashFlow) (*api.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"PostWriteOff",
	)
	defer span.Finish()

	err := s.operationsDB.ModifyOperation(
		ctx,
		in.GetUserId(),
		in.GetServiceId(),
		in.GetOrderId(),
		in.GetAmount(),
		approveCode,
	)

	if err != nil {
		s.logger.Error("cannot ModifyOperation", zap.Error(err))
		return &api.Empty{}, errors.Wrap(err, "cannot ModifyOperation")
	}

	return &api.Empty{}, nil
}

func (s server) PostGetReport(ctx context.Context, in *api.User) (*api.Report, error) {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"PostGetReport",
	)
	defer span.Finish()

	balanceHistory, err := s.bankAccountDB.GetBalanceHistory(ctx, in.GetUserId())

	if err != nil {
		s.logger.Error("cannot GetBalanceHistory", zap.Error(err))
		return &api.Report{}, errors.Wrap(err, "cannot GetBalanceHistory")
	}

	var resultHistory api.Report
	for _, operation := range balanceHistory {
		resultHistory.History = append(resultHistory.History, &api.BalanceHistory{
			Date:   timestamppb.New(operation.Date),
			Amount: operation.Amount,
		})
	}

	return &resultHistory, nil
}

func Run(config *config.Service, logger *zap.Logger, bankAccountDB bankAccountDB, operationsDB operationsDB) {
	grpcAddress := fmt.Sprintf("%s:%d", config.GetHost(), config.GetGrpcPort())
	grpcListener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		logger.Fatal("failed to listen grpc: %v", zap.Error(err))
	}

	serv := server{
		logger:        logger,
		bankAccountDB: bankAccountDB,
		operationsDB:  operationsDB,
	}

	s := grpc.NewServer()
	api.RegisterQueryListenerServer(s, &serv)

	ctx := context.Background()
	rmux := runtime.NewServeMux()
	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	{
		err := api.RegisterQueryListenerHandlerServer(ctx, rmux, serv)
		if err != nil {
			log.Fatal(err)
		}
	}

	httpAddress := fmt.Sprintf("%s:%d", config.GetHost(), config.GetHttpPort())
	httpListener, err := net.Listen("tcp", httpAddress)
	if err != nil {
		logger.Fatal("failed to listen http: %v", zap.Error(err))
	}

	go func() {
		reflection.Register(s)
		if err := s.Serve(grpcListener); err != nil {
			logger.Fatal("failed to serve: %v", zap.Error(err))
		}
	}()

	logger.Sugar().Infof("Serving http address %d", config.GetGrpcPort())
	err = http.Serve(httpListener, mux)
	if err != nil {
		logger.Fatal("failed to http serve", zap.Error(err))
	}
}
