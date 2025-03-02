package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// EstimatedPositions  テーブルのデータ構造体
type EstimatedPosition struct {
	ID           string `db:"id"`
	X            int    `db:"x"`
	Y            int    `db:"y"`
	CreatedAt    string `db:"created_at"`
	TrajectoryID string `db:"trajectory_id"`
}

func GetEstimatedPositionsByTrajectoryID(db *sql.DB, trajectoryID string) ([]EstimatedPosition, error) {

	// EstimatedPositionsテーブルの id,x,y,created_at,trajectory_idのデータ取得
	rows, err := db.Query("SELECT id,x,y,created_at,trajectory_id FROM estimated_positions WHERE trajectory_id = '01JET1DV4WJ2EP78B8HAKK5SP0'")
	if err != nil {
		return nil, fmt.Errorf("クエリ実行エラー %w", err)
	}
	defer rows.Close()

	// 結果をスライスに格納
	var estimatedPositions []EstimatedPosition
	for rows.Next() {
		var estimatedPosition EstimatedPosition
		if err := rows.Scan(&estimatedPosition.ID, &estimatedPosition.X, &estimatedPosition.Y, &estimatedPosition.CreatedAt, &estimatedPosition.TrajectoryID); err != nil {
			return nil, fmt.Errorf("データスキャンエラー: %w", err)
		}
		estimatedPositions = append(estimatedPositions, estimatedPosition)
	}

	return estimatedPositions, nil
}
