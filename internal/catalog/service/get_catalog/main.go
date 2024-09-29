package get_catalog

import (
	"encoding/json"
	"net/http"

	catalog_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog/domain"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/aws/s3_client"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewLogger() {
	internalLogger = logger.NewLoggerWithPrefix("getCatalogService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(ownerId *string) (*catalog_domain.Catalog, *util.PatternError) {
	body, err := s3_client.ReaderJson(ownerId)

	if err != nil {
		internalLogger.Error("Ocurred error when try read json", zap.NamedError("error", err))
		return nil, &util.PatternError{
			Code:         http.StatusNotFound,
			MessageError: err.Error(),
		}
	}

	var data catalog_domain.Catalog
	if err = json.Unmarshal(body, &data); err != nil {
		internalLogger.Error("Failed to unmarshal json", zap.NamedError("error", err))
		return nil, &util.PatternError{
			Code:         http.StatusBadRequest,
			MessageError: err.Error(),
		}
	}

	return &data, nil
}
