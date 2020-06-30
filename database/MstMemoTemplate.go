package dto

import (
	"database/sql"
	"time"
)

type MstMemoTemplate struct {
	TemplateID int            `db:"template_id"`
	MemoName   sql.NullString `db:"memo_name"`
	MemoHeader sql.NullString `db:"memo_header"`
	MemoFooter sql.NullString `db:"memo_footer"`
	CreatedBy  sql.NullString `db:"created_by"`
	UpdatedBy  sql.NullString `db:"updated_by"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
}
