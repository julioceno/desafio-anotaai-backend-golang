package sqs_client

import (
	"fmt"
	"lambda/aws/aws_session"
	"lambda/util"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Envs struct {
	catalogEmitConsumer string
}

var (
	svc                         *sqs.SQS
	QueueUrlCatalogEmitConsumer *string
)

func NewHandler() {
	envs := getEnvs()
	svc = sqs.New(aws_session.AwsSession)

	QueueUrlCatalogEmitConsumer = getQueueUrl(envs.catalogEmitConsumer)
}

func getQueueUrl(queueName string) *string {
	fmt.Println("Getting url service")
	url, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})

	if err != nil {
		log.Fatal("Ocurred error when try get url queue")
	}

	return url.QueueUrl
}

func getEnvs() Envs {
	catalogEmitConsumer := os.Getenv("CATALOG_EMIT_CONSUMER")
	util.ThrowErrorIfEnvNotExists("CATALOG_EMIT_CONSUMER", catalogEmitConsumer)

	return Envs{
		catalogEmitConsumer: catalogEmitConsumer,
	}
}
