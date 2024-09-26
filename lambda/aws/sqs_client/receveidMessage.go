package sqs_client

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ReceveiMessage(queueUrl *string) []*sqs.Message {
	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            queueUrl,
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(30),
	})

	if err != nil {
		fmt.Printf("Ocurred error when try get messages queue: %v", err)
	}

	return msgResult.Messages
}
