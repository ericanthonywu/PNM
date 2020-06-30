package dto

import (
	"database/sql"
	"time"
)

type MstSuratType struct {
	TypeID    int            `db:"type_id"`
	TypeName  sql.NullString `db:"type_name"`
	CreatedBy sql.NullString `db:"created_by"`
	UpdatedBy sql.NullString `db:"updated_by"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}
