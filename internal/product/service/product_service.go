package product_service

import (
	product_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/domain"
	product_repository "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/repository"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/product/service/create_product"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/product/service/delete_product"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/product/service/fetch_product"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/product/service/get_product"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/product/service/update_product"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
)

type ProductService struct{}

func NewProductService() *ProductService {
	product_repository.NewProductRepository()

	create_product.NewLogger()
	update_product.NewLogger()
	get_product.NewLogger()
	delete_product.NewLogger()
	fetch_product.NewLogger()

	return &ProductService{}
}

func (cs *ProductService) Create(data product_domain.CreateProduct) (*product_domain.Product, *util.PatternError) {
	return create_product.Run(data)
}

func (cs *ProductService) Update(id *string, data product_domain.UpdateProduct) (*product_domain.Product, *util.PatternError) {
	return update_product.Run(id, data)
}

func (cs *ProductService) Getproduct(id *string) (*product_domain.Product, *util.PatternError) {
	return get_product.Run(id)
}

func (cs *ProductService) Deleteproduct(id *string) *util.PatternError {
	return delete_product.Run(id)
}

func (cs *ProductService) Fetchproduct(filter *product_domain.Product, sortQuery *util.SortQuery) (*util.ResponseFormat, *util.PatternError) {
	return fetch_product.Run(filter, sortQuery)
}
