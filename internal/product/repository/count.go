package product_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *ProductRepository) Count(ctxMongo context.Context, filter primitive.D) (int64, error) {
	count, err := r.collection.CountDocuments(ctxMongo, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
