package update_product

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	product_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/domain"
	product_repository "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/repository"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

// TODO: adicionar validação pra verificar se a categoria que o produto ta sendo inserido, realmente existe
func NewLogger() {
	internalLogger = logger.NewLoggerWithPrefix("updateProductService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(id *string, data product_domain.UpdateProduct) (*product_domain.Product, *util.PatternError) {
	internalLogger.Info("Updating product...")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productToUpdate, err := getOldProduct(id, &ctxMongo)
	if err != nil {
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}

	updateFields(&data, productToUpdate)
	productUpdated, err := product_repository.Repository.Update(id, ctxMongo, productToUpdate)
	if err != nil {
		internalLogger.Error("Ocurred error when try update product", zap.NamedError("error", err))
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}

	internalLogger.Info("Product updated")
	return productUpdated, nil
}

func getOldProduct(id *string, ctxMongo *context.Context) (*product_domain.Product, error) {
	internalLogger.Info("Searching product...")
	productObtained := product_repository.Repository.FindById(id, *ctxMongo)

	if productObtained == nil {
		err := errors.New("product not exists")
		internalLogger.Error("Product not exists", zap.NamedError("error", err))
	}

	logger.Info("Product obtained")
	return productObtained, nil
}

func updateFields(data *product_domain.UpdateProduct, oldProduct *product_domain.Product) {
	if data.Title != nil {
		oldProduct.Title = *data.Title
	}

	if data.CategoryId != nil {
		oldProduct.CategoryId = *data.CategoryId
	}

	if data.Price != nil {
		oldProduct.Price = *data.Price
	}

	if data.Description != nil {
		oldProduct.Description = *data.Description
	}
}
