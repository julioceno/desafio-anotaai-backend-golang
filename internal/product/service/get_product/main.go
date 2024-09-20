package get_product

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

func NewLogger() {
	internalLogger = logger.NewLoggerWithPrefix("getProductServer")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(id *string) (*product_domain.Product, *util.PatternError) {
	internalLogger.Info("Getting product...")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	product := product_repository.Repository.FindById(id, ctxMongo)
	if product == nil {
		internalLogger.Error("Ocurred error when try get product")
		error := errors.New("product not exists")
		return nil, &util.PatternError{
			Code:         http.StatusNotFound,
			MessageError: error.Error(),
		}
	}

	internalLogger.Info("Product obtained")
	return product, nil
}
