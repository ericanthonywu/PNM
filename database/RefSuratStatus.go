package dto

import (
	"database/sql"
	"time"
)

type RefSuratStatus struct {
	SuratStatusID int            `db:"surat_status_id"`
	StatusCode    sql.NullString `db:"status_code"`
	Status        sql.NullString `db:"status"`
	CreatedBy     sql.NullString `db:"created_by"`
	UpdatedBy     sql.NullString `db:"updated_by"`
	CreatedAt     time.Time      `db:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at"`
}
