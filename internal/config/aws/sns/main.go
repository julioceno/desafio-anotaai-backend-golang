package sns_client

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/aws/aws_session"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"go.uber.org/zap"
)

var (
	svc            *sns.SNS
	internalLogger *zap.Logger

	CATALOG_EMITER = ""
)

func NewHandler() {
	internalLogger = logger.NewLoggerWithPrefix("sns")
	svc = sns.New(aws_session.AwsSession)
}
