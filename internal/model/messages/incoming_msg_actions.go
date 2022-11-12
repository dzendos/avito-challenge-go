package messages

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dzendos/avito-challenge/internal/types"
	"github.com/pkg/errors"
)

func (s *Model) credit(ctx context.Context, msg *Message) error {
	input := strings.Split(msg.Text, " ")
	enteredAmount, err := strconv.ParseInt(input[1], 10, 64)

	if err != nil {
		return errors.Wrap(err, "cannot ParseInt")
	}

	currencyRate, err := s.ratesDB.GetCurrencyRate(ctx, GetUserCurrency(msg.UserID), time.Now())

	amount := enteredAmount * int64(currencyRate) / 100

	reqBody, err := json.Marshal(map[string]string{
		"amount":  fmt.Sprintf("%d", amount),
		"user_id": fmt.Sprintf("%d", msg.UserID),
	})
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/api/credit", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return s.tgClient.SendMessage(err.Error(), msg.UserID)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "cannot io.ReadAll")
	}

	return s.tgClient.SendMessage(string(bytes), msg.UserID)
}

func (s *Model) reserve(ctx context.Context, msg *Message) error {
	orderID, serviceID, enteredAmount, err := getQueryInfo(msg.Text)

	if err != nil {
		return errors.Wrap(err, "cannot getQueryInfo")
	}

	currencyRate, err := s.ratesDB.GetCurrencyRate(ctx, GetUserCurrency(msg.UserID), time.Now())

	amount := enteredAmount * int64(currencyRate) / 100

	reqBody, err := json.Marshal(map[string]string{
		"user_id":    fmt.Sprintf("%d", msg.UserID),
		"order_id":   fmt.Sprintf("%d", orderID),
		"service_id": fmt.Sprintf("%d", serviceID),
		"amount":     fmt.Sprintf("%d", amount),
	})
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/api/reserve", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return s.tgClient.SendMessage(err.Error(), msg.UserID)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "cannot io.ReadAll")
	}

	return s.tgClient.SendMessage(string(bytes), msg.UserID)
}

func (s *Model) cancel(ctx context.Context, msg *Message) error {
	orderID, serviceID, enteredAmount, err := getQueryInfo(msg.Text)

	if err != nil {
		return errors.Wrap(err, "cannot getQueryInfo")
	}

	currencyRate, err := s.ratesDB.GetCurrencyRate(ctx, GetUserCurrency(msg.UserID), time.Now())

	amount := enteredAmount * int64(currencyRate) / 100

	reqBody, err := json.Marshal(map[string]string{
		"user_id":    fmt.Sprintf("%d", msg.UserID),
		"order_id":   fmt.Sprintf("%d", orderID),
		"service_id": fmt.Sprintf("%d", serviceID),
		"amount":     fmt.Sprintf("%d", amount),
	})
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/api/cancelreserve", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return s.tgClient.SendMessage(err.Error(), msg.UserID)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "cannot io.ReadAll")
	}

	return s.tgClient.SendMessage(string(bytes), msg.UserID)
}

func (s *Model) writeoff(ctx context.Context, msg *Message) error {
	orderID, serviceID, enteredAmount, err := getQueryInfo(msg.Text)

	if err != nil {
		return errors.Wrap(err, "cannot getQueryInfo")
	}

	currencyRate, err := s.ratesDB.GetCurrencyRate(ctx, GetUserCurrency(msg.UserID), time.Now())

	amount := enteredAmount * int64(currencyRate) / 100
	reqBody, err := json.Marshal(map[string]string{
		"user_id":    fmt.Sprintf("%d", msg.UserID),
		"order_id":   fmt.Sprintf("%d", orderID),
		"service_id": fmt.Sprintf("%d", serviceID),
		"amount":     fmt.Sprintf("%d", amount),
	})
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/api/writeoff", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return s.tgClient.SendMessage(err.Error(), msg.UserID)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "cannot io.ReadAll")
	}

	return s.tgClient.SendMessage(string(bytes), msg.UserID)
}

func (s *Model) getBalance(ctx context.Context, msg *Message) error {
	reqBody, err := json.Marshal(map[string]string{
		"user_id": fmt.Sprintf("%d", msg.UserID),
	})
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/api/getbalance", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return s.tgClient.SendMessage(err.Error(), msg.UserID)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "cannot io.ReadAll")
	}

	return s.tgClient.SendMessage(string(bytes), msg.UserID)
}

func (s *Model) getReport(ctx context.Context, msg *Message) error {
	reqBody, err := json.Marshal(map[string]string{
		"user_id": fmt.Sprintf("%d", msg.UserID),
	})
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/api/getreport", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return s.tgClient.SendMessage(err.Error(), msg.UserID)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "cannot io.ReadAll")
	}

	return s.tgClient.SendMessage(string(bytes), msg.UserID)
}

func (s *Model) changeCurrency(ctx context.Context, msg *Message) error {
	input := strings.Split(msg.Text, " ")

	enteredCurrency := strings.ToUpper(input[1])

	var currency types.Currency
	switch enteredCurrency {
	case "USD":
		currency = types.USD
	case "EUR":
		currency = types.EUR
	case "CNY":
		currency = types.CNY
	default:
		currency = types.RUB
	}

	SetUserCurrency(msg.UserID, currency)

	return s.tgClient.SendMessage("Текущая валюта: "+string(currency), msg.UserID)
}

func (s *Model) getCurrency(ctx context.Context, msg *Message) error {
	currency := GetUserCurrency(msg.UserID)

	return s.tgClient.SendMessage(string(currency), msg.UserID)
}

func getQueryInfo(text string) (int64, int64, int64, error) {
	input := strings.Split(text, " ")

	serviceID, err := strconv.ParseInt(input[1], 10, 64)
	if err != nil {
		return 0, 0, 0, errors.Wrap(err, "cannot ParseInt")
	}

	orderID, err := strconv.ParseInt(input[2], 10, 64)
	if err != nil {
		return 0, 0, 0, errors.Wrap(err, "cannot ParseInt")
	}

	amount, err := strconv.ParseInt(input[3], 10, 64)
	if err != nil {
		return 0, 0, 0, errors.Wrap(err, "cannot ParseInt")
	}

	return serviceID, orderID, amount, nil
}
