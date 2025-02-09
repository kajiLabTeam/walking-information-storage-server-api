package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// CorrectPositions  テーブルのデータ構造体
type CorrectPositions struct {
	ID           string `db:"id"`
	X            int    `db:"x"`
	Y            int    `db:"y"`
	CreatedAt    string `db:"created_at"`
	TrajectoryID string `db:"trajectory_id"`
}

func GetCorrectPositionsByTrajectoryID(db *sql.DB, trajectoryID Trajectories) (CorrectPositions, error) {
	var correctPosition CorrectPositions
	// correct_positions id,x,y,created_at,trajectory_idのデータ取得
	rows, err := db.Query("SELECT id,x,y,created_at,trajectory_id FROM correct_positions WHERE trajectory_id = '01JET1DV4WJ2EP78B8HAKK5SP0'")
	if err != nil {
		return CorrectPositions{}, fmt.Errorf("クエリ実行エラー %w", err)
	}
	defer rows.Close()

	// データ取得と出力
	for rows.Next() {

		if err := rows.Scan(&correctPosition.ID, &correctPosition.X, &correctPosition.Y, &correctPosition.CreatedAt, &correctPosition.TrajectoryID); err != nil {
			return CorrectPositions{}, fmt.Errorf("データスキャンエラー: %w", err)
		}
	}

	return CorrectPositions{ID: correctPosition.ID, X: correctPosition.X, Y: correctPosition.Y, CreatedAt: correctPosition.CreatedAt, TrajectoryID: correctPosition.TrajectoryID}, nil

}
