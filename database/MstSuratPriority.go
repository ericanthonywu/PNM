package dto

import (
	"database/sql"
	"time"
)

type MstSuratPriority struct {
	PriorityID   int            `db:"priority_id"`
	PriorityName sql.NullString `db:"priority_name"`
	CreatedBy    sql.NullString `db:"created_by"`
	UpdatedBy    sql.NullString `db:"updated_by"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}
