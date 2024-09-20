package category_repository

import (
	"context"

	category_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CategoryRepository) Create(ctxMongo context.Context, category category_domain.Category) (*category_domain.Category, error) {
	documentCreated, err := r.collection.InsertOne(ctxMongo, category)
	if err != nil {
		return nil, err
	}

	id := documentCreated.InsertedID.(primitive.ObjectID).Hex()
	categoryCreated := r.FindById(&id, ctxMongo)

	return categoryCreated, nil
}
