package category_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	OwnerId     string             `json:"ownerId" bson:"ownerId"`
}

type CreateCategory struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	OwnerId     string `json:"ownerId" validate:"required"`
}

type UpdateCategory struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description"`
}
