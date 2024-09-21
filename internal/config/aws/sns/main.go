package sns_client

import (
	"os"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/aws/aws_session"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

var (
	svc            *sns.SNS
	internalLogger *zap.Logger

	CATALOG_EMITER *string
)

func NewHandler() {
	internalLogger = logger.NewLoggerWithPrefix("sns")
	svc = sns.New(aws_session.AwsSession)

	catalogEmiter := os.Getenv("catalogEmiter_TOPIC")
	util.ThrowErrorIfEnvNotExists("CATALOG_EMITER_TOPIC", catalogEmiter)
	CATALOG_EMITER = &catalogEmiter
}
