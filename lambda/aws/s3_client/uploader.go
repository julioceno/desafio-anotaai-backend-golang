package s3_client

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploaderJson(bucketName *string, key *string, body interface{}) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	fmt.Println(*key)
	fmt.Println(body)
	jsonReader := bytes.NewReader(jsonData)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: bucketName,
		Key:    key,
		Body:   aws.ReadSeekCloser(jsonReader),
	})

	if err != nil {
		return err
	}

	return nil
}
