package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// floor_informationテーブルのデータ構造体
type FloorInformation struct {
	ID      string `db:"id"`
	floorID string `db:"floor_id"`
}

// Todo : floorIDを元にfloor_informationテーブルからフロア情報ID(floor_information_id)を取得
func GetFloorInformationIDByfloorID(db *sql.DB, floorID string) (FloorInformation, error) {
	var floorInformation FloorInformation

	return floorInformation, nil
}
