package delete_category

import (
	"context"
	"net/http"
	"time"

	catalog_emiter "github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog/service"
	category_repository "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/repository"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/category/service/get_category"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewLogger() {
	internalLogger = logger.NewLoggerWithPrefix("deleteCategoryService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(id *string) *util.PatternError {
	internalLogger.Info("Delete category...")

	internalLogger.Info("Getting current category to emiter catalog after...")
	currentCategory, _ := get_category.Run(id)

	ctxMongo, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := category_repository.Repository.Delete(id, ctxMongo)
	if err != nil {
		internalLogger.Error("Ocurred error when try delete category")
		return &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}

	catalog_emiter.Run(&currentCategory.OwnerId)
	internalLogger.Info("Category deleted")
	return nil
}

func callCatalogEmiter() {

}
