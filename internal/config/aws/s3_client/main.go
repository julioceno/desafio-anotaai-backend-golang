package s3_client

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/aws/aws_session"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/logger"
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/util"
	"go.uber.org/zap"
)

type Envs struct {
	catalogBucketName string
	awsEndpoints3     string
}

var (
	svc            *s3.S3
	internalLogger *zap.Logger
	downloader     *s3manager.Downloader

	CATALOG_BUCKET_NAME *string
)

func NewHandler() {
	internalLogger = logger.NewLoggerWithPrefix("s3")

	envs := getEnvs()
	svc = s3.New(aws_session.AwsSession, &aws.Config{
		Endpoint: aws.String(envs.awsEndpoints3),
	})
	downloader = s3manager.NewDownloader(aws_session.AwsSession)

	CATALOG_BUCKET_NAME = &envs.catalogBucketName
	_, err := svc.HeadBucket(&s3.HeadBucketInput{
		Bucket: CATALOG_BUCKET_NAME,
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
