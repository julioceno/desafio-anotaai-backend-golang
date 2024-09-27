package catalog

import (
	"context"
	"lambda/db"
	"slices"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Category struct {
	CategoryTitle       string    `json:"category_title"`
	CategoryDescription string    `json:"category_description"`
	Items               []Product `json:"items"`
}

type Catalog struct {
	OwnerId string     `json:"owner_id"`
	Catalog []Category `json:"catalog"`
}

func Create(ownerId *string) (*Catalog, error) {
	categories, products, err := getModels(ownerId)

	if err != nil {
		return nil, err
	}

	categoriesFormatted := []Category{}
	for _, category := range categories {
		categoryIdStr := category.Id.Hex()
		currentProducts := findAndRemoveProduct(&categoryIdStr, &products)

		category := Category{
			CategoryTitle:       category.Name,
			CategoryDescription: category.Description,
			Items:               currentProducts,
		}

		categoriesFormatted = append(categoriesFormatted, category)
	}

	catalog := Catalog{
		OwnerId: *ownerId,
		Catalog: categoriesFormatted,
	}

	return &catalog, nil
}

func getModels(ownerId *string) ([]db.Category, []db.Product, error) {
	filter := bson.D{
		bson.E{Key: "ownerId", Value: *ownerId},
	}
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var categories []db.Category
	var products []db.Product
	var err error

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		categories, err = getCategories(ctxMongo, filter)
	}()
	go func() {
		defer wg.Done()
		products, err = getProducts(ctxMongo, filter)
	}()

	wg.Wait()

	return categories, products, err
}

func getCategories(ctxMongo context.Context, filter primitive.D) ([]db.Category, error) {
	cursor, err := db.CategoryCollection.Find(ctxMongo, filter)
	if err != nil {
		return nil, err
	}

	var categories []db.Category
	for cursor.Next(ctxMongo) {
		var category db.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func getProducts(ctxMongo context.Context, filter primitive.D) ([]db.Product, error) {
	cursor, err := db.ProductCollection.Find(ctxMongo, filter)
	if err != nil {
		return nil, err
	}

	var products []db.Product
	for cursor.Next(ctxMongo) {
		var product db.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func findAndRemoveProduct(categoryId *string, products *[]db.Product) []Product {
	productsFromCategory := []Product{}

	for index, product := range *products {
		isThisCategory := product.CategoryId == *categoryId
		if !isThisCategory {
			continue
		}

		productFormatted := Product{
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
		}
		productsFromCategory = append(productsFromCategory, productFormatted)
		slices.Delete(*products, index, index+1)
		break
	}

	return productsFromCategory
}
