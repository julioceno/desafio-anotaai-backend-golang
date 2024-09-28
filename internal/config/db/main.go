package db

import (
	"context"
	"os"

	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Database *mongo.Client
)

func NewHandler() {
	logger.Info("Creating connection with database...")
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		logger.Fatal("Database url not exists", nil)
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))

	if err != nil {
		logger.Fatal("Occurred an error with mongo connection", err)
	}

	if err := client.Ping(ctx, options.Client().ReadPreference); err != nil {
		logger.Fatal("Occured an error in make ping in mongo connection", err)
	}

	logger.Info("Connection created")
	Database = client
}

func GetMongoCollection(db *mongo.Client, nameCollection string) *mongo.Collection {
	collection := db.Database("catalog").Collection(nameCollection)
	return collection
}
