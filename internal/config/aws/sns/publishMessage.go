package sns_client

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"go.uber.org/zap"
)

func PublishMessage(topicArn *string, message *string) {
	internalLogger.Info("Sending message")
	publishParams := &sns.PublishInput{
		Message:  message,
		TopicArn: topicArn,
	}

	_, err := svc.Publish(publishParams)
	if err != nil {
		internalLogger.Error("Ocurred error when send message", zap.NamedError("error", err))
	} else {
		internalLogger.Info("Message sent")
	}
}
