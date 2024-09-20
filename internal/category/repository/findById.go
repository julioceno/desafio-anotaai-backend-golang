package category_repository

import (
	"context"

	category_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/domain"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *CategoryRepository) FindById(id *string, ctxMongo context.Context) *category_domain.Category {
	objId, err := util.ConvertToObjectId(id)
	if err != nil {
		return nil
	}

	filter := bson.M{"_id": objId}
	document := r.collection.FindOne(ctxMongo, filter)

	var category category_domain.Category
	if err = document.Decode(&category); err != nil {
		return nil
	}

	return &category
}
