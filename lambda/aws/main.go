package aws_client

import (
	"lambda/aws/aws_session"
	"lambda/aws/s3_client"
	"lambda/aws/sqs_client"
)

func NewHandler() {
	aws_session.NewHandler()
	sqs_client.NewHandler()
	s3_client.NewHandler()
}
