package services

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Download(session *session.Session, floorID string, floorInformationID string) (string, error) {

	svc := s3.New(session)

	// 署名付きURLを作成するキー
	key := fmt.Sprintf("floors/%s/floor-information/%s/floor-map.png", floorID, floorInformationID)

	// 署名付きURLのリクエストを作成
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("indoor-location-estimation"),
		Key:    aws.String(key),
	})

	// リクエストの有効期限（例: 15分）
	expiration := 15 * time.Minute

	// 署名付きURLを取得
	presignedURL, err := req.Presign(expiration)
	if err != nil {
		log.Println("Error generating presigned URL:", err)
		return "", err
	}

	return presignedURL, err
}
