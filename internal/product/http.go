package product

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	product_domain "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/domain"
	product_service "github.com/julioceno/desafio-anotaai-backend-golang/internal/product/service"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewHandler(r *gin.Engine) {
	internalLogger = logger.NewLoggerWithPrefix("productController")
	productService := product_service.NewProductService()

	productRoutes := r.Group("products")
	productRoutes.POST("", func(ctx *gin.Context) {
		var body product_domain.CreateProduct
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

		response, patternError := productService.Create(body)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("POST - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, patternError.Code, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusOK, response)
	})

	productRoutes.PATCH("/:id", func(ctx *gin.Context) {
		id, err := util.GetIdParam(ctx)
		if err != nil {
			internalLogger.Error("PUT - Id not exists", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		var body product_domain.UpdateProduct
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

		response, patternError := productService.Update(&id, body)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("PUT - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, patternError.Code, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusOK, response)
	})

	productRoutes.GET("/:id", func(ctx *gin.Context) {
		id, err := util.GetIdParam(ctx)
		if err != nil {
			internalLogger.Error("GET/:id - Id not exists", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		response, patternError := productService.Getproduct(&id)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("GET/:id - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, patternError.Code, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusOK, response)
	})

	productRoutes.DELETE("/:id", func(ctx *gin.Context) {
		id, err := util.GetIdParam(ctx)
		if err != nil {
			internalLogger.Error("DELETE - Id not exists", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		patternError := productService.Deleteproduct(&id)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("DELETE - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, patternError.Code, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusNoContent, nil)
	})

	productRoutes.GET("", func(ctx *gin.Context) {
		filter, err := createFilter(ctx)
		if err != nil {
			internalLogger.Error("GET - Ocurred error when try get filter", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		sortOptions, err := createSortOptions(ctx)
		if err != nil {
			internalLogger.Error("GET - Ocurred error when try get sort options", zap.NamedError("error", err))
			util.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		response, patternError := productService.Fetchproduct(filter, sortOptions)
		if patternError != nil {
			errorBuilt := errors.New(patternError.MessageError)
			internalLogger.Error(fmt.Sprintf("GET - %v", "Ocurred error in service"), zap.NamedError("error", errorBuilt))
			util.SendError(ctx, patternError.Code, patternError.MessageError)
			return
		}

		util.SendSuccess(ctx, http.StatusOK, response)
	})
}

func createFilter(ctx *gin.Context) (*product_domain.Product, error) {
	priceFiltered := &product_domain.Product{
		Title:       ctx.Query("title"),
		CategoryId:  ctx.Query("categoryId"),
		Description: ctx.Query("description"),
		OwnerId:     ctx.Query("ownerId"),
	}

	priceStr := ctx.Query("price")
	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return nil, err
		}

		priceFiltered.Price = price
	}

	return priceFiltered, nil
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
