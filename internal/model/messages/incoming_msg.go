package messages

import (
	"context"
	"strings"
	"time"

	"github.com/dzendos/avito-challenge/internal/types"
)

type messageSender interface {
	SendMessage(text string, userID int64) error
	DeleteMessage(userID int64, messageID int) error
}

type currencyUpdater interface {
	UpdateCurrencyRate(ctx context.Context) error
}

type ratesDB interface {
	SetCurrencyRate(ctx context.Context, currency types.CurrencyRate) error
	GetCurrencyRate(ctx context.Context, currency types.Currency, date time.Time) (int64, error)
}

type Model struct {
	tgClient        messageSender
	ratesDB         ratesDB
	currencyUpdater currencyUpdater
}

func New(tgClient messageSender, updater currencyUpdater, ratesDB ratesDB) *Model {
	return &Model{
		tgClient:        tgClient,
		currencyUpdater: updater,
		ratesDB:         ratesDB,
	}
}

type Message struct {
	Text      string
	UserID    int64
	MessageID int
}

const (
	helpMsg = ""
)

func (s *Model) IncomingMessage(ctx context.Context, msg *Message) error {
	input := strings.Split(msg.Text, " ")
	switch input[0] {
	case "/start":
		return s.tgClient.SendMessage("hello", msg.UserID)
	case "/help":
		return s.tgClient.SendMessage(helpMsg, msg.UserID)
	case "/credit":
		return s.credit(ctx, msg)
	case "/reserve":
		return s.reserve(ctx, msg)
	case "/cancel":
		return s.cancel(ctx, msg)
	case "/writeoff":
		return s.writeoff(ctx, msg)
	case "/get_balance":
		return s.getBalance(ctx, msg)
	case "/get_report":
		return s.getReport(ctx, msg)
	case "/change_currency":
		return s.changeCurrency(ctx, msg)
	case "/get_currency":
		return s.getCurrency(ctx, msg)
	}

	return s.tgClient.SendMessage("не знаю эту команду", msg.UserID)
}
