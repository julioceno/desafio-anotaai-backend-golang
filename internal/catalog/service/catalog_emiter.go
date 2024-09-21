package catalog_emiter

import (
	"encoding/json"
	"fmt"

	sns_client "github.com/julioceno/desafio-anotaai-backend-golang/internal/config/aws/sns"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"go.uber.org/zap"
)

var (
	internalLogger *zap.Logger
)

type _messagePattern struct {
	OwnerId *string `json:"ownerId"`
}

func NewLogger() {
	internalLogger = logger.NewLoggerWithPrefix("catalogEmiterService")
	if internalLogger == nil {
		logger.Logger.Error("Not can get internal logger")
	}
}

func Run(ownerId *string) error {
	internalLogger.Info(fmt.Sprintf("Publish message in topic with owner id %s", *ownerId))
	message := _messagePattern{ownerId}

	jsonBody, err := json.Marshal(message)
	if err != nil {
		internalLogger.Error("Ocurred error when try convet message to json", zap.NamedError("error", err))
		return err
	}

	jsonFormatted := string(jsonBody)
	internalLogger.Info(jsonFormatted)

	sns_client.PublishMessage(sns_client.CATALOG_EMITER, &jsonFormatted)
	return nil
}
