package dto

import (
	"database/sql"
	"time"
)

type RelSuratRecipient struct {
	RecipientID   int            `db:"recipient_id"`
	RecipientName sql.NullString `db:"recipient_name"`
	RecipientNik  sql.NullString `db:"recipient_nik"`
	SuratID       sql.NullInt64  `db:"surat_id"`
	SignStatus    sql.NullInt64  `db:"sign_status"`
	Notes         sql.NullString `db:"notes"`
	CreatedBy     sql.NullString `db:"created_by"`
	UpdatedBy     sql.NullString `db:"updated_by"`
	CreatedAt     time.Time      `db:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at"`
}
