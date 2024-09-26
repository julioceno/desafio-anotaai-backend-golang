package s3_client

import (
	"fmt"
	"lambda/aws/aws_session"
	"lambda/util"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Envs struct {
	catalogBucketName string
	awsEndpoints3     string
}

var (
	svc                 *s3.S3
	CATALOG_BUCKET_NAME *string
)

func NewHandler() {
	envs := getEnvs()
	svc = s3.New(aws_session.AwsSession, &aws.Config{
		Endpoint: aws.String(envs.awsEndpoints3),
	})

	CATALOG_BUCKET_NAME = &envs.catalogBucketName

	_, err := svc.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String("catalog"),
	})

	if err != nil {
		fmt.Print("Não é possivel conectar ao bucekt")
		fmt.Println(err)
	}
}

func getEnvs() Envs {
	catalogBucketName := os.Getenv("CATALOG_BUCKET_NAME")
	util.ThrowErrorIfEnvNotExists("CATALOG_BUCKET_NAME", catalogBucketName)

	awsEndpoints3 := os.Getenv("AWS_ENDPOINT_S3")
	util.ThrowErrorIfEnvNotExists("AWS_ENDPOINT_S3", awsEndpoints3)

	return Envs{
		catalogBucketName: catalogBucketName,
		awsEndpoints3:     awsEndpoints3,
	}
}
