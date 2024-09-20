package product_repository

import (
	"context"

	product_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/domain"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *ProductRepository) FindById(id *string, ctxMongo context.Context) *product_domain.Product {
	objId, err := util.ConvertToObjectId(id)
	if err != nil {
		return nil
	}

	filter := bson.M{"_id": objId}
	document := r.collection.FindOne(ctxMongo, filter)

	var product product_domain.Product
	if err = document.Decode(&product); err != nil {
		return nil
	}

	return &product
}
