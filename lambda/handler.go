package main

import (
	"context"
	"encoding/json"
	"fmt"
	aws_client "lambda/aws"
	"lambda/aws/s3_client"
	"lambda/aws/sqs_client"
	"lambda/catalog"
	"lambda/db"
	"log"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/joho/godotenv"
)

type MessageWrapper struct {
	Type             string `json:"Type"`
	MessageId        string `json:"MessageId"`
	TopicArn         string `json:"TopicArn"`
	Message          string `json:"Message"`
	Timestamp        string `json:"Timestamp"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature        string `json:"Signature"`
	SigningCertURL   string `json:"SigningCertURL"`
}

type InnerMessage struct {
	OwnerId string `json:"ownerId"`
}

var wg sync.WaitGroup

func handler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	startDepencies()
	for {
		messages := sqs_client.ReceveiMessage(sqs_client.QueueUrlCatalogEmitConsumer)
		fmt.Printf("Consume %v messages \n", len(messages))
		for _, msg := range messages {
			wg.Add(1)

			decodedBody, err := getBodyMessage(msg.Body)
			if err != nil {
				return events.APIGatewayProxyResponse{}, err
			}

			go func() {
				defer wg.Done()
				createCatalog(&decodedBody.OwnerId, msg.ReceiptHandle)
			}()
		}
		wg.Wait()
	}
}

func startDepencies() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		db.NewHandler()
	}()

	go func() {
		defer wg.Done()
		aws_client.NewHandler()
	}()

	wg.Wait()
}

func createCatalog(ownerId *string, receiptHandle *string) {
	catalogBuilt, err := catalog.Create(ownerId)
	if err != nil {
		fmt.Printf("Ocurred error when try consume message from owner %s error: %v \n", *ownerId, err)
		return
	}

	if err := s3_client.UploaderJson(s3_client.CATALOG_BUCKET_NAME, ownerId, *catalogBuilt); err != nil {
		fmt.Printf("Is not possible post message in s3 from owner %s error: %v \n", *ownerId, err)
		return
	}

	fmt.Printf("Message consumed from owner %s, deleting message...\n", *ownerId)
	if err := sqs_client.DeleteMessage(sqs_client.QueueUrlCatalogEmitConsumer, receiptHandle); err != nil {
		fmt.Printf("ocurred error when try deleted message from owner %s, error: %v \n", *ownerId, err)
		return
	}
	fmt.Printf("Message deleted from owner %s \n", *ownerId)
}

func getBodyMessage(body *string) (*InnerMessage, error) {
	var wrapper MessageWrapper
	err := json.Unmarshal([]byte(*body), &wrapper)
	if err != nil {
		fmt.Printf("Error deserializing wrapper: %v\n", err)
		return nil, err
	}

	var innerMessage InnerMessage
	err = json.Unmarshal([]byte(wrapper.Message), &innerMessage)
	if err != nil {
		fmt.Printf("Error deserializing inner message: %v\n", err)
		return nil, err
	}

	fmt.Printf("Inner message deserialized successfully: %+v\n", innerMessage)
	return &innerMessage, nil
}
