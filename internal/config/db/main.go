package db

import (
	"context"
	"os"

	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	Database *mongo.Client
)

func NewHandler() {
	logger.Logger.Info("Creating connection with database...")
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		logger.Logger.Fatal("Database url not exists")
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))

	if err != nil {
		logger.Logger.Fatal("Occurred an error with mongo connection", zap.NamedError("error", err))
	}

	if err := client.Ping(ctx, options.Client().ReadPreference); err != nil {
		logger.Logger.Fatal("Occured an error in make ping in mongo connection", zap.NamedError("error", err))
	}

	logger.Logger.Info("Connection created")
	Database = client
}

func GetMongoCollection(db *mongo.Client, nameCollection string) *mongo.Collection {
	eventsCollection := db.Database("events").Collection(nameCollection)
	return eventsCollection
}
