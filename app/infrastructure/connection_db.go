package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// WalkingInformation テーブルのデータ構造体
type WalkingInformation struct {
	ID           string `db:"id"`
	PedestrianID string `db:"pedestrian_id"`
}

func ConnectionDB() (*sql.DB, error) {

	//環境変数の読み込み
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	// PostgreSQL 接続文字列の作成
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", user, password, name, host, port)
	fmt.Println(connStr)

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

	// walking_information テーブルのデータ取得
	rows, err := db.Query("SELECT id, pedestrian_id FROM walking_information")
	// fmt.Println(rows)

	if err != nil {
		return nil, fmt.Errorf("クエリ実行エラー %w", err)
	}
	defer rows.Close()

	fmt.Println(rows)
	// データ取得と出力
	for rows.Next() {

		fmt.Println("こんにゃく")
		var wi WalkingInformation
		if err := rows.Scan(&wi.ID, &wi.PedestrianID); err != nil {
			return nil, fmt.Errorf("データスキャンエラー: %w", err)
		}
		// fmt.Println("エラーは出てニアよ")
		fmt.Printf("ID: %s,PedestrianID :%s\n", wi.ID, wi.PedestrianID)
	}
	fmt.Println(rows)

	if err := db.Close(); err != nil {
		log.Printf("データベースのクローズに失敗: %v", err)
	}
	return db, nil

}
