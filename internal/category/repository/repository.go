package category_repository

import (
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

var (
	Repository *CategoryRepository
)

func NewCategoryRepository() {
	Repository = &CategoryRepository{
		collection: db.GetMongoCollection(db.Database, db.CollectionsName.CATEGORY),
	}
}
