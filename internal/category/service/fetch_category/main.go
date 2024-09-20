package fetch_category

import (
	"context"
	"net/http"
	"time"

	category_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/domain"
	category_repository "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/repository"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewLogger() {
	internalLogger = logger.NewLoggerWithPrefix("fetchCategoriesService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(filter *category_domain.Category, sortQuery *util.SortQuery) (*util.ResponseFormat, *util.PatternError) {
	filterBuilt := buildFilter(filter)
	sortOptionsBuilt := buildSortOptions(sortQuery)

	internalLogger.Info("Fetching categories...")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	categories, err := category_repository.Repository.Fetch(ctxMongo, filterBuilt, sortOptionsBuilt)
	if err != nil {
		internalLogger.Error("Ocurred error when try fetch categories")
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}
	internalLogger.Info("Categories obtained")

	count, err := category_repository.Repository.Count(ctxMongo, filterBuilt)
	if err != nil {
		internalLogger.Error("Ocurred error when try get count categories")
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}
	internalLogger.Info("Count obtained")

	response := util.ResponseFormat{
		Count: count,
		Data:  categories,
	}

	return &response, nil
}

func buildFilter(filter *category_domain.Category) primitive.D {
	filterBuilt := bson.D{
		{"name", primitive.Regex{Pattern: filter.Name, Options: "i"}},
	}

	if filter.OwnerId != "" {
		filterBuilt = append(filterBuilt, bson.E{"ownerId", filter.OwnerId})
	}

	return filterBuilt
}

func buildSortOptions(sortOptions *util.SortQuery) *options.FindOptions {
	opts := options.Find().SetSkip(int64(sortOptions.Skip)).SetLimit(int64(sortOptions.Limit))
	return opts
}
