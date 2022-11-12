package microservice

import (
	"context"
	"os"
	"os/signal"
	"testing"

	"github.com/dzendos/avito-challenge/api"
	"github.com/dzendos/avito-challenge/cmd/logging"
	"github.com/dzendos/avito-challenge/microservice/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_OnPostCredit_ShouldAddRefill(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := logging.InitLogger()
	bankAccountDB := mocks.NewMockbankAccountDB(ctrl)
	operationsDB := mocks.NewMockoperationsDB(ctrl)
	s := server{
		logger:        logger,
		bankAccountDB: bankAccountDB,
		operationsDB:  operationsDB,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	bankAccountDB.EXPECT().AddRefill(gomock.Any(), gomock.Any(), gomock.Any())

	_, err := s.PostCredit(ctx, &api.CreditInfo{
		UserId: 1230,
		Amount: 145,
	})

	assert.NoError(t, err)
}

func Test_OnPostGetBalance_ShouldGetAmount(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := logging.InitLogger()
	bankAccountDB := mocks.NewMockbankAccountDB(ctrl)
	operationsDB := mocks.NewMockoperationsDB(ctrl)
	s := server{
		logger:        logger,
		bankAccountDB: bankAccountDB,
		operationsDB:  operationsDB,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	bankAccountDB.EXPECT().GetAmount(gomock.Any(), gomock.Any())

	_, err := s.PostGetBalance(ctx, &api.User{
		UserId: 123,
	})

	assert.NoError(t, err)
}

func Test_OnPostReserve_ShouldReserve(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := logging.InitLogger()
	bankAccountDB := mocks.NewMockbankAccountDB(ctrl)
	operationsDB := mocks.NewMockoperationsDB(ctrl)
	s := server{
		logger:        logger,
		bankAccountDB: bankAccountDB,
		operationsDB:  operationsDB,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	bankAccountDB.EXPECT().GetAmount(gomock.Any(), gomock.Any())
	bankAccountDB.EXPECT().AddRefill(gomock.Any(), gomock.Any(), gomock.Any())
	operationsDB.EXPECT().AddOperation(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())

	_, err := s.PostReserve(ctx, &api.CashFlow{
		UserId:    123,
		ServiceId: 1,
		OrderId:   2,
		Amount:    0,
	})

	assert.NoError(t, err)
}

func Test_OnPostCancelReserve_ShouldModifyAndRefill(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := logging.InitLogger()
	bankAccountDB := mocks.NewMockbankAccountDB(ctrl)
	operationsDB := mocks.NewMockoperationsDB(ctrl)
	s := server{
		logger:        logger,
		bankAccountDB: bankAccountDB,
		operationsDB:  operationsDB,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	operationsDB.EXPECT().ModifyOperation(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
	bankAccountDB.EXPECT().AddRefill(gomock.Any(), gomock.Any(), gomock.Any())

	_, err := s.PostCancelReserve(ctx, &api.CashFlow{
		UserId:    123,
		ServiceId: 1,
		OrderId:   2,
		Amount:    0,
	})

	assert.NoError(t, err)
}

func Test_OnPostWriteoff_ShouldModify(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := logging.InitLogger()
	bankAccountDB := mocks.NewMockbankAccountDB(ctrl)
	operationsDB := mocks.NewMockoperationsDB(ctrl)
	s := server{
		logger:        logger,
		bankAccountDB: bankAccountDB,
		operationsDB:  operationsDB,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	operationsDB.EXPECT().ModifyOperation(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())

	_, err := s.PostWriteOff(ctx, &api.CashFlow{
		UserId:    123,
		ServiceId: 1,
		OrderId:   2,
		Amount:    0,
	})

	assert.NoError(t, err)
}

func Test_OnPostGetReport_ShouldGetBalanceHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := logging.InitLogger()
	bankAccountDB := mocks.NewMockbankAccountDB(ctrl)
	operationsDB := mocks.NewMockoperationsDB(ctrl)
	s := server{
		logger:        logger,
		bankAccountDB: bankAccountDB,
		operationsDB:  operationsDB,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	bankAccountDB.EXPECT().GetBalanceHistory(gomock.Any(), gomock.Any())

	_, err := s.PostGetReport(ctx, &api.User{
		UserId: 123,
	})

	assert.NoError(t, err)
}
