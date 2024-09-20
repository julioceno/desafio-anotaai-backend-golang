package category_repository

import (
	"context"

	category_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/domain"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *CategoryRepository) Update(id *string, ctxMongo context.Context, category *category_domain.Category) (*category_domain.Category, error) {
	objId, err := util.ConvertToObjectId(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	update := bson.M{
		"$set": category,
	}

	r.collection.FindOneAndUpdate(ctxMongo, filter, update)
	categoryUpdated := r.FindById(id, ctxMongo)
	return categoryUpdated, nil
}
