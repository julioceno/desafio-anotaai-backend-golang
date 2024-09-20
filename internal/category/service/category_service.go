package category_service

import (
	category_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/domain"
	category_repository "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/repository"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/category/service/create_category"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/category/service/delete_category"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/category/service/fetch_category"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/category/service/get_category"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/category/service/update_category"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
)

type CategoryService struct {}

func NewCategoryService() *CategoryService {
	category_repository.NewCategoryRepository()

	create_category.NewLogger()
	update_category.NewLogger()
	get_category.NewLogger()
	delete_category.NewLogger()
	fetch_category.NewLogger()

	return &CategoryService{}
}

func (cs *CategoryService) Create(data category_domain.CreateCategory) (*category_domain.Category, *util.PatternError) {
	return create_category.Run(data)
}

func (cs *CategoryService) Update(id *string, data category_domain.UpdateCategory) (*category_domain.Category, *util.PatternError) {
	return update_category.Run(id, data)
}

func (cs *CategoryService) GetCategory(id *string) (*category_domain.Category, *util.PatternError) {
	return get_category.Run(id)
}

func (cs *CategoryService) DeleteCategory(id *string) *util.PatternError {
	return delete_category.Run(id)
}

func (cs *CategoryService) FetchCategory(filter *category_domain.Category, sortQuery *util.SortQuery) (*util.ResponseFormat, *util.PatternError) {
	return fetch_category.Run(filter, sortQuery)
}
