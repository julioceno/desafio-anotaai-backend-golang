package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/category"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/product"
)

func NewHandler() {
	r := gin.Default()
	category.NewHandler(r)
	product.NewHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	logger.Info(fmt.Sprintf("Init server in port %s", port))
	r.Run(port)
}
