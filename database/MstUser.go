package dto

import (
	"database/sql"
	"time"
)

type MstUser struct {
	UserID           int            `db:"user_id"`
	Username         sql.NullString `db:"username"`
	Password         string         `db:"password"`
	Email            sql.NullString `db:"email"`
	Phone            sql.NullString `db:"phone"`
	Device           sql.NullString `db:"device"`
	Os               sql.NullString `db:"os"`
	DeviceToken      sql.NullString `db:"device_token"`
	VerifKey         sql.NullString `db:"verif_key"`
	VerifExpiredDate time.Time      `db:"verif_expired_date"`
	OtpNumber        sql.NullString `db:"otp_number"`
	Apikey           sql.NullString `db:"apikey"`
	LastLoginIp      sql.NullString `db:"last_login_ip"`
	IsBlocked        sql.NullInt64  `db:"is_blocked"`
	PositionID       sql.NullInt64  `db:"position_id"`
	DivisionID       sql.NullInt64  `db:"division_id"`
	CreatedBy        sql.NullString `db:"created_by"`
	UpdatedBy        sql.NullString `db:"updated_by"`
}
