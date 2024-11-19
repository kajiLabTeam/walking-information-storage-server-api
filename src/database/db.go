package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // PostgreSQL用ドライバ
)

// ConnectDBはPostgreSQLデータベースに接続します。
func SqlConnect() (*sql.DB, error) {
	// 環境変数から接続情報を取得
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// DSN (Data Source Name) の設定
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// データベース接続
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("データベース接続エラー:", err)
		return nil, err
	}

	// 接続確認
	if err := db.Ping(); err != nil {
		fmt.Println("データベース接続エラー:", err)
		return nil, err
	}

	fmt.Println("データベースに接続しました！")
	return db, nil
}
