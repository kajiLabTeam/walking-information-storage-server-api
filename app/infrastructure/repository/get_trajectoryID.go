package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Trajectories  テーブルのデータ構造体
type Trajectories struct {
	ID string `db:"id"`
}

func GetTrajectoryIDByFloorID(db *sql.DB, id string) (Trajectories, error) {

	// EstimatedPositionsテーブルの id,x,y,created_at,trajectory_idのデータ取得
	rows, err := db.Query("SELECT id  FROM trajectories WHERE floor_id = '01F8VYXK67BGC1F9RP1E4S9YTV'")
	if err != nil {
		return Trajectories{}, fmt.Errorf("クエリ実行エラー %w", err)
	}
	defer rows.Close()

	var trajectory Trajectories
	// データ取得と出力
	for rows.Next() {

		if err := rows.Scan(&trajectory.ID); err != nil {
			return Trajectories{}, fmt.Errorf("データスキャンエラー: %w", err)
		}

	}

	return Trajectories{ID: trajectory.ID}, nil

}
