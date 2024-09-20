package fetch_product

import (
	"context"
	"net/http"
	"time"

	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	product_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/domain"
	product_repository "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/repository"
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
	internalLogger = logger.NewLoggerWithPrefix("fetchProductsService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(filter *product_domain.Product, sortQuery *util.SortQuery) (*util.ResponseFormat, *util.PatternError) {
	filterBuilt := buildFilter(filter)
	sortOptionsBuilt := buildSortOptions(sortQuery)

	internalLogger.Info("Fetching products...")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	products, err := product_repository.Repository.Fetch(ctxMongo, filterBuilt, sortOptionsBuilt)
	if err != nil {
		internalLogger.Error("Ocurred error when try fetch products")
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}
	internalLogger.Info("Products obtained")

	count, err := product_repository.Repository.Count(ctxMongo, filterBuilt)
	if err != nil {
		internalLogger.Error("Ocurred error when try get count products")
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}
	internalLogger.Info("Count obtained")

	response := util.ResponseFormat{
		Count: count,
		Data:  products,
	}

	return &response, nil
}

func buildFilter(filter *product_domain.Product) primitive.D {
	filterBuilt := bson.D{
		{"title", primitive.Regex{Pattern: filter.Title, Options: "i"}},
		{"description", primitive.Regex{Pattern: filter.Description, Options: "i"}},
	}

	if filter.OwnerId != "" {
		filterBuilt = append(filterBuilt, bson.E{"ownerId", filter.OwnerId})
	}

	if filter.CategoryId != "" {
		filterBuilt = append(filterBuilt, bson.E{"categoryId", filter.CategoryId})
	}

	if filter.Price > 0 {
		filterBuilt = append(filterBuilt, bson.E{"price", filter.Price})
	}

	return filterBuilt
}

func buildSortOptions(sortOptions *util.SortQuery) *options.FindOptions {
	opts := options.Find().SetSkip(int64(sortOptions.Skip)).SetLimit(int64(sortOptions.Limit))
	return opts
}
