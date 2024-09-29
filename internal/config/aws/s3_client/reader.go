package s3_client

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.uber.org/zap"
)

func ReaderJson(ownerId *string) ([]byte, error) {
	key := fmt.Sprintf("%s.json", *ownerId)

	out, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: CATALOG_BUCKET_NAME,
		Key:    aws.String(key),
	})

	if err != nil {
		internalLogger.Error("Occurred error when reading json", zap.NamedError("error", err))
		return nil, err
	}
	defer out.Body.Close()

	body, err := ioutil.ReadAll(out.Body)
	if err != nil {
		internalLogger.Error("Failed to read body", zap.NamedError("error", err))
		return nil, err
	}

	return body, nil
}
