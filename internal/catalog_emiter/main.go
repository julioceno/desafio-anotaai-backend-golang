package catalog

import (
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/catalog_emiter"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

func NewHandler() {
	catalog_emiter.NewHandler()
}
