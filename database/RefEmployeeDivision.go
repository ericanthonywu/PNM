package dto

import (
	"database/sql"
	"time"
)

type RefEmployeeDivision struct {
	DivisionID   int            `db:"division_id"`
	DivisionCode sql.NullString `db:"division_code"`
	DivisionName sql.NullString `db:"division_name"`
	CreatedBy    sql.NullString `db:"created_by"`
	UpdatedBy    sql.NullString `db:"updated_by"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}
