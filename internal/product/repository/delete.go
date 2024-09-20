package product_repository

import (
	"context"
	"errors"

	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *ProductRepository) Delete(id *string, ctxMongo context.Context) error {
	objId, err := util.ConvertToObjectId(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objId}
	deletedResult, err := r.collection.DeleteOne(ctxMongo, filter)
	if !(deletedResult.DeletedCount > 0) {
		error := errors.New("no deleted any document")
		return error
	}

	return nil
}
