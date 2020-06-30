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

type RelSuratRecipientMobile struct {
	RecipientID   int    `db:"recipient_id"`
	RecipientName string `db:"recipient_name"`
	RecipientNik  string `db:"recipient_nik"`
	SuratID       int    `db:"surat_id"`
	SignStatus    int    `db:"sign_status"`
	Notes         string `db:"notes"`
	CreatedBy     string `db:"created_by"`
	UpdatedBy     string `db:"updated_by"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
}
