package worker

import (
	"context"
	"log"

	"github.com/dzendos/avito-challenge/internal/model/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type updateFetcher interface {
	Start() tgbotapi.UpdatesChannel
	Request(callback tgbotapi.CallbackConfig) error
	Stop()
}

type MessageHandler interface {
	IncomingMessage(ctx context.Context, msg *messages.Message) error
}

type updateListenerWorker struct {
	updateFetcher  updateFetcher
	messageHandler MessageHandler
}

func NewUpdateListenerWorker(updateFetcher updateFetcher,
	messageHandler MessageHandler) *updateListenerWorker {
	return &updateListenerWorker{
		updateFetcher:  updateFetcher,
		messageHandler: messageHandler,
	}
}

func (w *updateListenerWorker) Run(ctx context.Context) {
	updates := w.updateFetcher.Start()

	for {
		select {
		case <-ctx.Done():
			w.updateFetcher.Stop()
			return
		case update, ok := <-updates:
			if !ok {
				w.updateFetcher.Stop()
				return
			}
			err := w.HandleUpdate(ctx, update)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (w *updateListenerWorker) HandleUpdate(ctx context.Context, update tgbotapi.Update) error {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		err := w.messageHandler.IncomingMessage(ctx, &messages.Message{
			Text:      update.Message.Text,
			UserID:    update.Message.From.ID,
			MessageID: update.Message.MessageID,
		})

		if err != nil {
			return errors.Wrap(err, "cannot IncomingMessage")
		}
	}

	return nil
}
