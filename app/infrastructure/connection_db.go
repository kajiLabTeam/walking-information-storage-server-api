package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// trajectories  テーブルのデータ構造体
type Trajectories struct {
	ID string `db:"id"`
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

	// trajectories テーブルのtrajectoriesID のデータ取得
	rows, err := db.Query("SELECT id  FROM trajectories WHERE floor_id = '01F8VYXK67BGC1F9RP1E4S9YTV'")

	if err != nil {
		return nil, fmt.Errorf("クエリ実行エラー %w", err)
	}
	defer rows.Close()

	// データ取得と出力
	for rows.Next() {

		var trajectory Trajectories
		if err := rows.Scan(&trajectory.ID); err != nil {
			return nil, fmt.Errorf("データスキャンエラー: %w", err)
		}

		fmt.Printf("trajectoryID: %s\n", trajectory.ID)
	}
	fmt.Println(rows)

	if err := db.Close(); err != nil {
		log.Printf("データベースのクローズに失敗: %v", err)
	}
	return db, nil

}
