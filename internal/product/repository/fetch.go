package product_repository

import (
	"context"

	product_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *ProductRepository) Fetch(ctxMongo context.Context, filter primitive.D, sortOptions *options.FindOptions) ([]product_domain.Product, error) {
	cursor, err := r.collection.Find(ctxMongo, filter, sortOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctxMongo)

	var products []product_domain.Product
	for cursor.Next(ctxMongo) {
		var product product_domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
