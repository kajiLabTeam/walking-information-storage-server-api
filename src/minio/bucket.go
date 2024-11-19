package minio

import (
	"fmt"

	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Bucket はバケットを表す構造体です。
func Getbuckets() (*s3.ListBucketsOutput, error) {
	// バケットの一覧を取得
	result, err := minioclient.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to list buckets, %v", err)
	}
	return result, nil
}
