package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/kajiLabTeam/walking-information-storage-server-api/app/infrastructure/repository"
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

	//floor_id(フロアID)を元に軌跡ID(trajectry_id)を取得
	repository.GetTrajectoryIDByFloorID(db, "01F8VYXK67BGC1F9RP1E4S9YTV")

	//軌跡ID(trajectry_id)に紐付いた推定座標(estimated_positions)を取得
	repository.GetEstimatedPositionsByTrajectoryID(db, "01JET1DV4WJ2EP78B8HAKK5SP0")

	if err := db.Close(); err != nil {
		log.Printf("データベースのクローズに失敗: %v", err)
	}

	return db, nil

}
