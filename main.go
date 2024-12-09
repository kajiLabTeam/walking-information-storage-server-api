package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {
	connStr := "user=xr-project-root password=9Q(/TnL8PY3d dbname=indoor_location_estimation sslmode=disable host=localhost port=5432"

	//データベースと接続
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	//defer:処理が完了次に実行
	defer db.Close()

	fmt.Println("DB接続")

	//SQL
	//検索結果を格納
	var pedestrianID string
	id := "01JBJR8PMCEA0KVAAM0X5R5PNJ"
	err = db.QueryRow("SELECT pedestrian_id FROM walking_information WHERE id=$1", id).Scan(&pedestrianID)
	if err != nil {
		log.Fatalln("歩行者ID取得できない!!")
	}
	fmt.Println(pedestrianID)

	//Minio
	//環境変数を読み込み
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	minioRootUser := os.Getenv("MINIO_ROOT_USER")
	minioRootPassword := os.Getenv("MINIO_ROOT_PASSWORD")
	minioRegion := os.Getenv("MINIO_REGION")
	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
	minioSecretKey := os.Getenv("MINIO_SECRET_KEY")
	minioBucketName := os.Getenv("MINIO_BUCKET_NAME")

	fmt.Println(minioRootUser)
	fmt.Println(minioRootPassword)
	fmt.Println(minioRegion)
	fmt.Println(minioEndpoint)
	fmt.Println(minioAccessKey)
	fmt.Println(minioSecretKey)
	fmt.Println(minioBucketName)
	fmt.Println()

	//sessionに接続
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(minioRegion),
		Credentials:      credentials.NewStaticCredentials(minioAccessKey, minioSecretKey, ""),
		Endpoint:         aws.String(minioEndpoint),
		S3ForcePathStyle: aws.Bool(true),
	})

	if err != nil {
		log.Fatal(err)
	}

	svc := s3.New(sess)

	fmt.Println(svc)

	// バケット一覧出力
	fmt.Println("バケット")
	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, bucket := range result.Buckets {
		fmt.Printf("%s\n", aws.StringValue(bucket.Name))
	}

	//ファイル取得  バケットとキー
	fmt.Println(minioBucketName)
	obj, err := svc.GetObject(&s3.GetObjectInput{

		Bucket: aws.String("indoor-location-estimation"),
		//取ってきたいファイル
		Key: aws.String("gyroscope/01J5AEZX8WCH456MAR137WPV7M/01J5AF8GPCP5BHSFYF02AX3VD0.csv"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(obj)
	// CSVリーダーを作成する
	reader := csv.NewReader(obj.Body)

	// CSVデータを全て読み込む
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV data:", err)
		return
	}
	// 読み込んだデータを表示する
	for _, record := range records {
		fmt.Println(record)
	}
}

//変数宣言

// var str string = "Hello,World!"
// fmt.Println(str)

// str1 := "Hello"
// fmt.Println(str1)
