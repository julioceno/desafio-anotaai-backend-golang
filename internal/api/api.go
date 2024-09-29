package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	catalog "github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/category"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/product"
)

func NewHandler() {
	r := gin.Default()
	category.NewHandler(r)
	product.NewHandler(r)
	catalog.NewHandler(r)

	port := getPort()
	logger.Info(fmt.Sprintf("Init server in port %s", port))
	r.Run(port)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}

	return ":" + port
}
