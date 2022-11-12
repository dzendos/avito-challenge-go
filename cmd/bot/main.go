package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/dzendos/avito-challenge/internal/clients/tg"
	"github.com/dzendos/avito-challenge/internal/config"
	"github.com/dzendos/avito-challenge/internal/currency"
	"github.com/dzendos/avito-challenge/internal/database"
	"github.com/dzendos/avito-challenge/internal/model/messages"
	"github.com/dzendos/avito-challenge/internal/worker"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	fmt.Println("initializing config")
	config, err := config.New()
	if err != nil {
		fmt.Println("config init failed:", err)
	}

	fmt.Println("initializing database")
	db, err := database.New(config)
	if err != nil {
		fmt.Println("database init failed", err)
	}

	ratesDB := database.NewRatesDB(db)

	fmt.Println("initializing telegram client")
	tgClient, err := tg.New(config)
	if err != nil {
		fmt.Println("tg client init failed:", err)
	}

	currencyUpdateModel := currency.NewCbrCurrencyUpdater(config, ratesDB)

	msgModel := messages.New(tgClient, currencyUpdateModel, ratesDB)

	currencyRateWorker := worker.NewCurrencyRateWorker(currencyUpdateModel)
	updateListenerWorker := worker.NewUpdateListenerWorker(tgClient, msgModel)

	currencyRateWorker.Run(ctx, config.GetUpdateRate())
	updateListenerWorker.Run(ctx)
}
