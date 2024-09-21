package util

import (
	"fmt"

	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
)

func ThrowErrorIfEnvNotExists(key string, value string) {
	if value == "" {
		logger.Fatal(fmt.Sprintf("%s não existe", key), nil)
	}
}
