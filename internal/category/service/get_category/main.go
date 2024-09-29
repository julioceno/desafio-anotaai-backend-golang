package get_category

import (
	"context"
	"errors"
	"net/http"
	"time"

	category_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/domain"
	category_repository "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/repository"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewLogger() {
	internalLogger = logger.NewLoggerWithPrefix("getCategoryService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(id *string) (*category_domain.Category, *util.PatternError) {
	internalLogger.Info("Getting category...")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	category := category_repository.Repository.FindById(id, ctxMongo)
	if category == nil {
		internalLogger.Error("Ocurred error when try get category")
		error := errors.New("category not exists")
		return nil, &util.PatternError{
			Code:         http.StatusNotFound,
			MessageError: error.Error(),
		}
	}

	internalLogger.Info("Category obtained")
	return category, nil
}
