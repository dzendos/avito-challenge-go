package main

import (
	"github.com/dzendos/avito-challenge/cmd/logging"
	"github.com/dzendos/avito-challenge/cmd/tracing"
	"github.com/dzendos/avito-challenge/internal/config"
	"github.com/dzendos/avito-challenge/internal/database"
	"github.com/dzendos/avito-challenge/microservice"
	"go.uber.org/zap"
)

func main() {
	logger := logging.InitLogger()
	tracing.InitTracing("query_handler", logger)

	logger.Info("initializing config")
	config, err := config.New()
	if err != nil {
		logger.Fatal("config init failed:", zap.Error(err))
	}

	logger.Info("initializing database")
	db, err := database.New(config)
	if err != nil {
		logger.Fatal("database init failed", zap.Error(err))
	}

	bankAccountDB := database.NewBankAccountDB(db)
	operationsDB := database.NewOperationsDB(db)

	microservice.Run(config, logger, bankAccountDB, operationsDB)
}
