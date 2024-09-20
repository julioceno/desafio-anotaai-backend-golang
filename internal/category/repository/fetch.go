package category_repository

import (
	"context"

	category_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *CategoryRepository) Fetch(ctxMongo context.Context, filter primitive.D, sortOptions *options.FindOptions) ([]category_domain.Category, error) {
	cursor, err := r.collection.Find(ctxMongo, filter, sortOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctxMongo)

	var categories []category_domain.Category
	for cursor.Next(ctxMongo) {
		var category category_domain.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
