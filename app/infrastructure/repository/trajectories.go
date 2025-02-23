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

func GetTrajectoryIDByFloorID(db *sql.DB, floorID string) ([]*Trajectories, error) {

	// EstimatedPositionsテーブルの id,x,y,created_at,trajectory_idのデータ取得
	//SQLクエリ
	query := "SELECT id  FROM trajectories WHERE floor_id = $1"
	rows, err := db.Query(query, floorID)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行エラー: %w", err)
	}
	defer rows.Close()

	var trajectories []*Trajectories

	// データ取得とスライスへの格納
	for rows.Next() {
		var t Trajectories
		if err := rows.Scan(&t.ID); err != nil {
			return nil, fmt.Errorf("データスキャンエラー: %w", err)
		}
		trajectories = append(trajectories, &t)
	}
	fmt.Println(trajectories)

	// ループ中のエラー確認
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("行の取得エラー: %w", err)
	}

	return trajectories, nil

}
