package minio

import (
	"bytes"
	"fmt"

	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListObject(bucketName string) (*s3.ListObjectsOutput, error) {
	result, err := minioclient.ListObjects(&s3.ListObjectsInput{
		Bucket: &bucketName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list objects, %v", err)
	}
	return result, nil
}

func GetObject(bucketName string, objectName string) (string, error) {
    result, err := minioclient.GetObject(&s3.GetObjectInput{
        Bucket: &bucketName,
        Key:    &objectName,
    })
    if err != nil {
        return "", fmt.Errorf("failed to get object, %v", err)
    }

    // コンテンツを読み取る
    buf := new(bytes.Buffer)
    _, err = buf.ReadFrom(result.Body)
    if err != nil {
        return "", fmt.Errorf("failed to read object content, %v", err)
    }

    return buf.String(), nil
}


