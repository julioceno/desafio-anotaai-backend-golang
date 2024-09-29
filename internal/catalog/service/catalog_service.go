package catalog_service

import (
	catalog_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog/domain"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog/service/catalog_emit"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog/service/get_catalog"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
)

type _CatalogService struct{}

var (
	Service *_CatalogService
)

func NewCategoryService() *_CatalogService {
	catalog_emit.NewLogger()
	get_catalog.NewLogger()

	Service = &_CatalogService{}
	return Service
}

func (cs *_CatalogService) Create(ownerId *string) error {
	return catalog_emit.Run(ownerId)
}

func (cs *_CatalogService) GetCatalog(ownerId *string) (*catalog_domain.Catalog, *util.PatternError) {
	return get_catalog.Run(ownerId)
}
