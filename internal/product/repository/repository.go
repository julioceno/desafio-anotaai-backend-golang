package product_repository

import (
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

var (
	Repository *ProductRepository
)

func NewProductRepository() {
	Repository = &ProductRepository{
		collection: db.GetMongoCollection(db.Database, db.CollectionsName.PRODUCT),
	}
}
