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

func GeCorrectPositionsByTrajectoryID(db *sql.DB, id string) (CorrectPositions, error) {

	// correct_positions id,x,y,created_at,trajectory_idのデータ取得
	rows, err := db.Query("SELECT id,x,y,created_at,trajectory_id FROM correct_positions WHERE trajectory_id = '01JET1DV4WJ2EP78B8HAKK5SP0'")
	if err != nil {
		return CorrectPositions{}, fmt.Errorf("クエリ実行エラー %w", err)
	}
	defer rows.Close()

	// データ取得と出力
	for rows.Next() {
		var correctPosition CorrectPositions
		if err := rows.Scan(&correctPosition.ID, &correctPosition.X, &correctPosition.Y, &correctPosition.CreatedAt, &correctPosition.TrajectoryID); err != nil {
			return CorrectPositions{}, fmt.Errorf("データスキャンエラー: %w", err)
		}

		fmt.Printf("correctPosition:ID %s,x %d,y %d ,CreatedAt %s,TrajectoryID %s \n", correctPosition.ID, correctPosition.X, correctPosition.Y, correctPosition.CreatedAt, correctPosition.TrajectoryID)
	}

	return CorrectPositions{}, nil

}
