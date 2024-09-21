package delete_product

import (
	"context"
	"net/http"
	"time"

	catalog_emiter "github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog_emiter/service"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	product_repository "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/repository"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/product/service/get_product"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewLogger() {
	internalLogger = logger.NewLoggerWithPrefix("deleteProductServer")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(id *string) *util.PatternError {
	internalLogger.Info("Delete product...")

	internalLogger.Info("Getting current product to emiter catalog after...")
	currentProduct, _ := get_product.Run(id)

	ctxMongo, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := product_repository.Repository.Delete(id, ctxMongo)
	if err != nil {
		internalLogger.Error("Ocurred error when try delete product")
		return &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}

	internalLogger.Info("Calling catalog emitter...")
	catalog_emiter.Run(&currentProduct.OwnerId)
	internalLogger.Info("Product deleted")
	return nil
}
