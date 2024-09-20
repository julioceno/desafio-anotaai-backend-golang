package update_category

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
	internalLogger = logger.NewLoggerWithPrefix("updateCategoryService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(id *string, data category_domain.UpdateCategory) (*category_domain.Category, *util.PatternError) {
	internalLogger.Info("Updating category...")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	categoryToUpdate, err := getOldCategory(id, &ctxMongo)
	if err != nil {
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}

	updateFields(data, categoryToUpdate)
	categoryUpdated, err := category_repository.Repository.Update(id, ctxMongo, categoryToUpdate)
	if err != nil {
		internalLogger.Error("Ocurred error when try update category", zap.NamedError("error", err))
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}

	internalLogger.Info("Category updated")
	return categoryUpdated, nil
}


func getOldCategory(id *string, ctxMongo *context.Context) (*category_domain.Category, error) {
	internalLogger.Info("Searching category...")
	categoryObtained := category_repository.Repository.FindById(id, *ctxMongo)

	if categoryObtained == nil {
		err := errors.New("category not exists")
		internalLogger.Error("Category not exists", zap.NamedError("error", err))
		return nil, err
	}

	logger.Info("Category obtained")
	return categoryObtained, nil
}

func updateFields(data category_domain.UpdateCategory, oldCategory *category_domain.Category) {
	if data.Name != nil {
		oldCategory.Name = *data.Name
	}

	if data.Description != nil {
		oldCategory.Description = *data.Description
	}
}
