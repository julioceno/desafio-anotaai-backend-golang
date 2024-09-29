package catlog

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	catalog_service "github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog/service"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewHandler(r *gin.Engine) {
	internalLogger = logger.NewLoggerWithPrefix("catalogController")
	catalogService := catalog_service.NewCategoryService()

	catalogRoutes := r.Group("catalogs")
	catalogRoutes.GET("/:ownerId", func(ctx *gin.Context) {
		ownerId, err := util.GetValueByParams(ctx, "ownerId")
		if err != nil {
			internalLogger.Error("GET/:ownerId - Owner not exists", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		response, patternError := catalogService.GetCatalog(&ownerId)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("GET/:ownerId - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, patternError.Code, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusOK, response)
	})
}
