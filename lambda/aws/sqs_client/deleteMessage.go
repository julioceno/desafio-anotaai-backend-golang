package sqs_client

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

func DeleteMessage(queueUrl *string, receiptHandle *string) error {
	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueUrl,
		ReceiptHandle: receiptHandle,
	})

	return err
}
