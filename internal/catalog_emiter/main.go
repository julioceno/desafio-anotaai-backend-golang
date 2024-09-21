package catalog

import (
	catalog_emiter "github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog_emiter/service"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewHandler() {
	catalog_emiter.NewLogger()
}
