package minio

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var minioclient = ConnectMinio()

func ConnectMinio() *s3.S3 {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	if endpoint == "" || accessKey == "" || secretKey == "" {
		log.Fatalf("環境変数が設定されていません。MINIO_ENDPOINT=%s, MINIO_ACCESS_KEY=%s, MINIO_SECRET_KEY=%s", endpoint, accessKey, secretKey)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Endpoint:         aws.String(endpoint),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		DisableSSL:       aws.Bool(false),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("AWSセッションの作成に失敗しました: %v", err)
	}

	svc := s3.New(sess)
	log.Println("Minioへの接続に成功しました")
	return svc
}