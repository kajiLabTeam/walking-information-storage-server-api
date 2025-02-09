package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// EstimatedPositions  テーブルのデータ構造体
type EstimatedPositions struct {
	ID           string `db:"id"`
	X            int    `db:"x"`
	Y            int    `db:"y"`
	CreatedAt    string `db:"created_at"`
	TrajectoryID string `db:"trajectory_id"`
}

func GetEstimatedPositionsByTrajectoryID(db *sql.DB, id string) (EstimatedPositions, error) {

	// EstimatedPositionsテーブルの id,x,y,created_at,trajectory_idのデータ取得
	rows, err := db.Query("SELECT id,x,y,created_at,trajectory_id FROM estimated_positions WHERE trajectory_id = '01JET1DV4WJ2EP78B8HAKK5SP0'")
	if err != nil {
		return EstimatedPositions{}, fmt.Errorf("クエリ実行エラー %w", err)
	}
	defer rows.Close()

	// データ取得と出力
	for rows.Next() {
		var estimatedPosition EstimatedPositions
		if err := rows.Scan(&estimatedPosition.ID, &estimatedPosition.X, &estimatedPosition.Y, &estimatedPosition.CreatedAt, &estimatedPosition.TrajectoryID); err != nil {
			return EstimatedPositions{}, fmt.Errorf("データスキャンエラー: %w", err)
		}

		fmt.Printf("estimated_position:ID %s,x %d,y %d ,CreatedAt %s,TrajectoryID %s \n", estimatedPosition.ID, estimatedPosition.X, estimatedPosition.Y, estimatedPosition.CreatedAt, estimatedPosition.TrajectoryID)
	}

	return EstimatedPositions{}, nil

}
