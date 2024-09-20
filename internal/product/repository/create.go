package product_repository

import (
	"context"

	product_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *ProductRepository) Create(ctxMongo context.Context, product product_domain.Product) (*product_domain.Product, error) {
	documentCreated, err := r.collection.InsertOne(ctxMongo, product)
	if err != nil {
		return nil, err
	}

	id := documentCreated.InsertedID.(primitive.ObjectID).Hex()
	productCreated := r.FindById(&id, ctxMongo)
	return productCreated, nil
}
