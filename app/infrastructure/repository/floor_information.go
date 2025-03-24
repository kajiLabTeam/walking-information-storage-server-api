package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// floor_informationテーブルのデータ構造体
type FloorInformation struct {
	ID        string `db:"id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	FloorID   string `db:"floor_id"`
}

// Todo : floorIDを元にfloor_informationテーブルからフロア情報ID(floor_information_id)を取得
func GetFloorInformationIDByfloorID(db *sql.DB, floorID string) (*FloorInformation, error) {

	fmt.Println("GetFloorInformationIDByfloorID", floorID)
	var floorInformation FloorInformation

	query := "SELECT id, created_at, updated_at, floor_id FROM floor_information WHERE floor_id = $1"
	rows, err := db.Query(query, floorID)

	if err != nil {
		return nil, fmt.Errorf("クエリ実行エラー %w", err)
	}
	defer rows.Close()

	// 結果をスライスに格納
	for rows.Next() {

		if err := rows.Scan(&floorInformation.ID, &floorInformation.CreatedAt, &floorInformation.UpdatedAt, &floorInformation.FloorID); err != nil {
			return nil, fmt.Errorf("データスキャンエラー: %w", err)
		}
	}

	return &floorInformation, nil
}
