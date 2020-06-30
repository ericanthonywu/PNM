package dto

import (
	"time"
)

type MstSuratStatus struct {
	StatusID    int       `db:"status_id"`
	StatusCode  string    `db:"status_code"`
	Status      string    `db:"status"`
	StatusAlias string    `db:"status_alias"`
	CreatedBy   uint64    `db:"created_by"`
	CreatedDate time.Time `db:"created_date"`
}
