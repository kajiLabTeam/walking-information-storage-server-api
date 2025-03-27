package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/lib/pq"
)

func ConnectionDB() (*sql.DB, error) {

	//環境変数の読み込み
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	// PostgreSQL 接続文字列の作成
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", user, password, name, host, port)

	//データベースと接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("データベース接続エラー: %w", err)
	}

	// 接続確認(Pingを飛ばす)
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("データベースの応答なし: %w", err)
	}
	fmt.Println("DB接続成功")

	// if err := db.Close(); err != nil {
	// 	log.Printf("データベースのクローズに失敗: %v", err)
	// }

	return db, nil

}

func ConnectionMinio() (*session.Session, error) {

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

	// sessionに接続
	session, err := session.NewSession(&aws.Config{
		Region:           aws.String(minioRegion),
		Credentials:      credentials.NewStaticCredentials(minioAccessKey, minioSecretKey, ""),
		Endpoint:         aws.String(minioEndpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("Minio接続エラー: %w", err)
	}

	return session, nil

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// svc := s3.New(sess)

	// fmt.Println(svc)

	// // バケット一覧出力
	// fmt.Println("バケット")
	// result, err := svc.ListBuckets(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, bucket := range result.Buckets {
	// 	fmt.Printf("%s\n", aws.StringValue(bucket.Name))
	// }

	// // ファイル取得  バケットとキー
	// fmt.Println(minioBucketName)
	// obj, err := svc.GetObject(&s3.GetObjectInput{

	// 	Bucket: aws.String("indoor-location-estimation"),
	// 	//取ってきたいファイル
	// 	Key: aws.String("floors/${floor_id}/floor-information/${floor_information_id}/floor-map.png"),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(obj)
	// // CSVリーダーを作成する
	// reader := csv.NewReader(obj.Body)

	// // CSVデータを全て読み込む
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	fmt.Println("Error reading CSV data:", err)
	// 	return
	// }
	// // 読み込んだデータを表示する
	// for _, record := range records {
	// 	fmt.Println(record)
	// }

}
