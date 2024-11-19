package model

import "time"

// Pedestrian は pedestrians テーブルに対応する構造体です。
type Pedestrian struct {
	ID        string     `json:"id" db:"id"`                           // 主キー
	CreatedAt time.Time  `json:"created_at" db:"created_at"`           // 作成日時（タイムゾーン付き）
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`           // 更新日時（タイムゾーン付き）
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"` // 削除日時（タイムゾーン付き, null許容）
}
