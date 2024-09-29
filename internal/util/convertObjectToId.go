package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertToObjectId(id *string) (*primitive.ObjectID, error) {
	objId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}

	return &objId, nil
}
