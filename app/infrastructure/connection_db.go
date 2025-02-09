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
	trajectory, err := repository.GetTrajectoryIDByFloorID(db, "01F8VYXK67BGC1F9RP1E4S9YTV")
	if err != nil {
		return nil, fmt.Errorf("軌跡IDの取得エラー: %w", err)
	}
	fmt.Println("trajectory.ID")

	fmt.Println(trajectory.ID)

	//軌跡ID(trajectry_id)に紐付いた推定座標(estimated_positions)/正解座標(correct_positions)を取得
	estimatedPositions, err := repository.GetEstimatedPositionsByTrajectoryID(db, trajectory.ID)
	if err != nil {
		return nil, fmt.Errorf("推定座標の取得エラー: %w", err)
	}

	correctPositions, err := repository.GetCorrectPositionsByTrajectoryID(db, trajectory.ID)
	if err != nil {
		return nil, fmt.Errorf("正解座標の取得エラー: %w", err)
	}

	//表示
	fmt.Println(estimatedPositions)
	fmt.Println(correctPositions)

	if err := db.Close(); err != nil {
		log.Printf("データベースのクローズに失敗: %v", err)
	}

	return db, nil

}
