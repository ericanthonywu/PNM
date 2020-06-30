package dto

import (
	"database/sql"
	"time"
)

type LogActivity struct {
	LogID        int            `db:"log_id"`
	UserID       sql.NullInt64  `db:"user_id"`
	Username     sql.NullString `db:"username"`
	ModuleAccess sql.NullString `db:"module_access"`
	Activity     sql.NullString `db:"activity"`
	OldData      sql.NullString `db:"old_data"`
	NewData      sql.NullString `db:"new_data"`
	CreatedBy    sql.NullString `db:"created_by"`
	UpdatedBy    sql.NullString `db:"updated_by"`
	CreatedAt    time.Time      `db:"created_at"`
}
