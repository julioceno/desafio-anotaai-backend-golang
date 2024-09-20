package category

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	category_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/domain"
	category_service "github.com/julioceno/desafio-anotaai-backend-golang/internal/category/service"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewHandler(r *gin.Engine) {
	internalLogger = logger.NewLoggerWithPrefix("categoryController")
	categoryService := category_service.NewCategoryService()

	categoriesRoutes := r.Group("categories")
	categoriesRoutes.POST("", func(ctx *gin.Context) {
		var body category_domain.CreateCategory
		if err := util.DecodeBody(ctx, &body); err != nil {
			internalLogger.Error("POST - Ocurred error when try decode body", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		if err := util.ValidateBody(body); err != nil {
			internalLogger.Error("POST - Ocurred error when validate body", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		response, patternError := categoryService.Create(body)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("POST - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, patternError.Code, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusOK, response)
	})

	categoriesRoutes.PATCH("/:id", func(ctx *gin.Context) {
		id, err := util.GetIdParam(ctx)
		if err != nil {
			internalLogger.Error("PUT - Id not exists", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		var body category_domain.UpdateCategory
		if err := util.DecodeBody(ctx, &body); err != nil {
			internalLogger.Error("PUT - Ocurred error when try decode body", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		if err := util.ValidateBody(body); err != nil {
			internalLogger.Error("PUT - Ocurred error when validate body", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		response, patternError := categoryService.Update(&id, body)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("PUT - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, http.StatusBadRequest, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusOK, response)
	})

	categoriesRoutes.GET("/:id", func(ctx *gin.Context) {
		id, err := util.GetIdParam(ctx)
		if err != nil {
			internalLogger.Error("GET/:id - Id not exists", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		response, patternError := categoryService.GetCategory(&id)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("GET/:id - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, patternError.Code, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusOK, response)
	})

	categoriesRoutes.DELETE("/:id", func(ctx *gin.Context) {
		id, err := util.GetIdParam(ctx)
		if err != nil {
			internalLogger.Error("DELETE - Id not exists", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		patternError := categoryService.DeleteCategory(&id)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("DELETE - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, patternError.Code, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusNoContent, nil)
	})

	categoriesRoutes.GET("", func(ctx *gin.Context) {
		filter := createFilter(ctx)
		sortOptions, err := createSortOptions(ctx)
		if err != nil {
			internalLogger.Error("GET - Ocurred error when try get sort options", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		response, patternError := categoryService.FetchCategory(filter, sortOptions)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("GET - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, patternError.Code, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusOK, response)
	})

}

func createFilter(ctx *gin.Context) *category_domain.Category {
	return &category_domain.Category{
		Name:    ctx.Query("name"),
		OwnerId: ctx.Query("ownerId"),
	}
}

func createSortOptions(ctx *gin.Context) (*util.SortQuery, error) {
	skip := ctx.DefaultQuery("skip", "0")
	limit := ctx.DefaultQuery("limit", "10")

	skipBase64, err := util.ConvertToNumber(skip)
	if err != nil {
		return nil, err
	}

	limitBase64, err := util.ConvertToNumber(limit)
	if err != nil {
		return nil, err
	}

	return &util.SortQuery{
		Skip:  int(skipBase64),
		Limit: int(limitBase64),
	}, nil
}
