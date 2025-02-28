package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

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
