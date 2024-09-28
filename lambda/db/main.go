package db

import (
	"context"
	"fmt"
	"lambda/util"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Category struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	OwnerId     string             `json:"ownerId" bson:"ownerId"`
}

type Product struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	CategoryId  string             `json:"categoryId" bson:"categoryId"`
	Price       float64            `json:"price" bson:"price"`
	Description string             `json:"description" bson:"description"`

	OwnerId string `json:"ownerId" bson:"ownerId"`
}

var (
	CategoryCollection *mongo.Collection
	ProductCollection  *mongo.Collection
	categoryName       = "categories"
	productName        = "products"
)

func NewHandler() {
	fmt.Println("Creating connection with database...")

	databaseUrl := os.Getenv("DATABASE_URL")
	util.ThrowErrorIfEnvNotExists("DATABASE_URL", databaseUrl)

	ctx := context.Background()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseUrl))
	if err != nil {
		log.Fatal("Occurred an error with mongo connection", err)
	}

	if err := db.Ping(ctx, options.Client().ReadPreference); err != nil {
		log.Fatal("Occured an error in make ping in mongo connection", err)
	}

	fmt.Println("Connection created, connecting with models...")
	CategoryCollection = getMongoCollection(db, categoryName)
	ProductCollection = getMongoCollection(db, productName)
	fmt.Println("Created models")
}

func getMongoCollection(db *mongo.Client, nameCollection string) *mongo.Collection {
	eventsCollection := db.Database("catalog").Collection(nameCollection)
	return eventsCollection
}
