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

func GetTrajectoriesIDByFloorID(db *sql.DB, floorID string) ([]Trajectory, error) {

	// EstimatedPositionsテーブルの id,x,y,created_at,trajectory_idのデータ取得
	//SQLクエリ
	query := "SELECT id  FROM trajectories WHERE floor_id = $1"
	rows, err := db.Query(query, floorID)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行エラー: %w", err)
	}
	defer rows.Close()

	// 結果をスライスに格納
	var trajectories []Trajectory
	for rows.Next() {
		var trajectory Trajectory
		if err := rows.Scan(&trajectory.ID); err != nil {
			return nil, fmt.Errorf("データスキャンエラー: %w", err)
		}
		trajectories = append(trajectories, trajectory)
	}

	return trajectories, nil

}
