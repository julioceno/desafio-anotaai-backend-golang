package create_category

import (
	"context"
	"net/http"
	"time"

	catalog_emiter "github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog/service"
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
	internalLogger = logger.NewLoggerWithPrefix("createCategoryService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(data category_domain.CreateCategory) (*category_domain.Category, *util.PatternError) {
	internalLogger.Info("Creating category...")

	categoryToCreate := category_domain.Category{
		Name:        data.Name,
		OwnerId:     data.OwnerId,
		Description: data.Description,
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	categoryCreated, err := category_repository.Repository.Create(ctxMongo, categoryToCreate)
	if err != nil {
		internalLogger.Error("Ocurred error when try create category", zap.NamedError("error", err))
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}

	catalog_emiter.Run(&categoryCreated.OwnerId)
	internalLogger.Info("Category Created")
	return categoryCreated, nil
}
