package product_repository

import (
	"context"

	product_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/domain"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *ProductRepository) Update(id *string, ctxMongo context.Context, product *product_domain.Product) (*product_domain.Product, error) {
	objId, err := util.ConvertToObjectId(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	update := bson.M{
		"$set": product,
	}

	r.collection.FindOneAndUpdate(ctxMongo, filter, update)
	productCreated := r.FindById(id, ctxMongo)
	return productCreated, nil
}
