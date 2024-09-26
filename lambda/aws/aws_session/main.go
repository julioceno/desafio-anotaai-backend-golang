package aws_session

import (
	"fmt"
	"lambda/util"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Envs struct {
	awsUrl          string
	awsRegion       string
	awsAccessKeyID  string
	awsSecretKey    string
	awsSessionToken string
}

var (
	AwsSession *session.Session
)

func NewHandler() {
	envs := getEnvs()

	fmt.Println("Creating aws session")
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:    aws.String(envs.awsUrl),
		Region:      aws.String(envs.awsRegion),
		Credentials: credentials.NewStaticCredentials(envs.awsAccessKeyID, envs.awsSecretKey, envs.awsSessionToken),
	}))
	fmt.Println("Session created, create queue instance")

	AwsSession = sess
}

func getEnvs() Envs {
	awsUrl := os.Getenv("AWS_URL")
	util.ThrowErrorIfEnvNotExists("AWS_URL", awsUrl)

	awsRegion := os.Getenv("AWS_REGION")
	util.ThrowErrorIfEnvNotExists("AWS_REGION", awsRegion)

	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	util.ThrowErrorIfEnvNotExists("AWS_ACCESS_KEY_ID", awsAccessKeyID)

	awsSecretKey := os.Getenv("AWS_SECRET_KEY")
	util.ThrowErrorIfEnvNotExists("AWS_SECRET_KEY", awsSecretKey)

	awsSessionToken := os.Getenv("AWS_SESSION_TOKEN")
	util.ThrowErrorIfEnvNotExists("AWS_SESSION_TOKEN", awsSessionToken)

	return Envs{
		awsUrl:          awsUrl,
		awsRegion:       awsRegion,
		awsAccessKeyID:  awsAccessKeyID,
		awsSecretKey:    awsSecretKey,
		awsSessionToken: awsSessionToken,
	}
}
