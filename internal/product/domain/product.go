package product_domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	CategoryId  string             `json:"categoryId" bson:"categoryId"`
	Price       float64            `json:"price" bson:"price"`
	Description string             `json:"description" bson:"description"`

	OwnerId string `json:"ownerId" bson:"ownerId"`
}

type CreateProduct struct {
	Title       string  `json:"title" validate:"required"`
	CategoryId  string  `json:"categoryId" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`

	OwnerId string `json:"ownerId" bson:"ownerId"`
}

type UpdateProduct struct {
	Title       *string  `json:"title"`
	CategoryId  *string  `json:"categoryId"`
	Price       *float64 `json:"price"`
	Description *string  `json:"description"`
}
