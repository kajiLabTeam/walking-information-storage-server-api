package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Trajectories  テーブルのデータ構造体

type Trajectory struct {
	ID string `db:"id"`
}

func GetTrajectoryIDByFloorID(db *sql.DB, floorID string) (*Trajectory, error) {

	// EstimatedPositionsテーブルの id,x,y,created_at,trajectory_idのデータ取得
	query := "SELECT id  FROM trajectories WHERE floor_id = $1"
	rows := db.QueryRow(query, floorID)

	var trajectory Trajectory

	// データ取得と出力

	if err := rows.Scan(&trajectory.ID); err != nil {

		return nil, fmt.Errorf("データスキャンエラー: %w", err)
	}

	return &Trajectory{ID: trajectory.ID}, nil

}
