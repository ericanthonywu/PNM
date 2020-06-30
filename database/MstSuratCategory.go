package dto

import (
	"database/sql"
	"time"
)

type MstSuratCategory struct {
	CategoryID   int            `db:"category_id"`
	CategoryName sql.NullString `db:"category_name"`
	CreatedBy    sql.NullString `db:"created_by"`
	UpdatedBy    sql.NullString `db:"updated_by"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}
