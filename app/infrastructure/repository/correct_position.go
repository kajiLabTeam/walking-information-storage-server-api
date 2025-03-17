package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// CorrectPositions  テーブルのデータ構造体
type CorrectPosition struct {
	ID           string `db:"id"`
	X            int    `db:"x"`
	Y            int    `db:"y"`
	CreatedAt    string `db:"created_at"`
	TrajectoryID string `db:"trajectory_id"`
}

func GetCorrectPositionsByTrajectoryID(db *sql.DB, trajectoryID string) ([]CorrectPosition, error) {
	// correct_positions id,x,y,created_at,trajectory_idのデータ取得

	query := "SELECT id,x,y,created_at,trajectory_id FROM correct_positions WHERE trajectory_id = $1"
	rows, err := db.Query(query, trajectoryID)

	if err != nil {
		return nil, fmt.Errorf("クエリ実行エラー %w", err)
	}
	defer rows.Close()

	// 結果をスライスに格納
	var correctPositions []CorrectPosition
	for rows.Next() {
		var correctPosition CorrectPosition
		if err := rows.Scan(&correctPosition.ID, &correctPosition.X, &correctPosition.Y, &correctPosition.CreatedAt, &correctPosition.TrajectoryID); err != nil {
			return nil, fmt.Errorf("データスキャンエラー: %w", err)
		}
		correctPositions = append(correctPositions, correctPosition)
	}

	return correctPositions, nil

}
