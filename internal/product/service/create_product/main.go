package create_product

import (
	"context"
	"net/http"
	"time"

	catalog_service "github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog/service"
	category_service "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/service"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	product_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/domain"
	product_repository "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/repository"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewLogger() {
	internalLogger = logger.NewLoggerWithPrefix("createProductService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(data product_domain.CreateProduct) (*product_domain.Product, *util.PatternError) {
	internalLogger.Info("Creating product...")
	categoryRecord, patternError := category_service.Service.GetCategory(&data.CategoryId)
	if patternError != nil {
		return nil, patternError
	}

	productToCreate := product_domain.Product{
		Title:       data.Title,
		CategoryId:  data.CategoryId,
		Price:       data.Price,
		Description: data.Description,
		OwnerId:     categoryRecord.OwnerId,
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productCreated, err := product_repository.Repository.Create(ctxMongo, productToCreate)
	if err != nil {
		internalLogger.Error("Ocurred error when try create product", zap.NamedError("error", err))
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}

	catalog_service.Service.Create(&productCreated.OwnerId)
	internalLogger.Info("Product Created")
	return productCreated, nil
}
