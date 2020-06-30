package dto

import (
	"database/sql"
	"time"
)

type RefEmployeePosition struct {
	PositionID   int            `db:"position_id"`
	PositionCode sql.NullString `db:"position_code"`
	PositionName sql.NullString `db:"position_name"`
	CreatedBy    sql.NullString `db:"created_by"`
	UpdatedBy    sql.NullString `db:"updated_by"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}
