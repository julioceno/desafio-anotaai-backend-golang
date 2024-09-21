package main

import (
	"github.com/joho/godotenv"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/api"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/aws"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/db"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"go.uber.org/zap"
)

func main() {
	logger.NewLogger()

	if err := godotenv.Load(); err != nil {
		logger.Logger.Fatal("Error loading .env file", zap.NamedError("error", err))
	}

	aws.NewHandler()
	db.NewHandler()
	api.NewHandler()
}
