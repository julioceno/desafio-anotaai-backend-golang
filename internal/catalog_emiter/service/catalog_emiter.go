package catalog_emiter

import (
	"encoding/json"

	sns_client "github.com/julioceno/desafio-anotaai-backend-golang/internal/config/aws/sns"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

type _messagePattern struct {
	ownerId *string
}

func NewLogger() {
	internalLogger = logger.NewLoggerWithPrefix("catalogEmiterService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(ownerId *string) error {
	internalLogger.Info("Publish message in topic")
	message := _messagePattern{ownerId}
	jsonBody, err := json.Marshal(message)

	if err != nil {
		internalLogger.Error("Ocurred error when try convet message to json", zap.NamedError("error", err))
		return err
	}

	jsonFormatted := string(jsonBody)
	sns_client.PublishMessage(sns_client.CATALOG_EMITER, &jsonFormatted)
	return nil
}
