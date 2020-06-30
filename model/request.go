package model

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type (
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	DefaultShowData struct {
		Limit  uint64 `json:"limit"`
		Offset uint64 `json:"offset"`
	}
	FindId struct {
		Id uint64 `json:"id"`
	}
	DefaultDelete struct {
		Id uint64 `json:"id"`
	}
	AddUser struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		IsBlocked bool   `json:"is_blocked"`
	}
	ToogleIsBlocked struct {
		IsBlocked int `json:"is_blocked"`
		Id        int `json:"id"`
	}
	MasterCrudDefault struct {
		Id   uint64 `json:"id"`
		Name string `json:"name"`
		// surat status req
		Code  string `json:"code"`
		Alias string `json:"alias"`
		// memo template
		Header string `json:"header"`
		Footer string `json:"footer"`
		// trx surat status
		StatusId   int    `json:"status_id"`
		StatusName string `json:"status_name"`
		// pagination
		Limit  uint64 `json:"limit"`
		Offset uint64 `json:"offset"`
		// assigned user
		SuratId int   `json:"surat_id"`
		UsersId []int `json:"users_id"`
	}
	SuratMasukReq struct {
		SuratID        uint64 `json:"surat_id"`
		Classification string `json:"classification"`
		CategoryID     uint64 `json:"category_id"`
		PriorityID     uint64 `json:"priority_id"`
		TypeID         uint64 `json:"type_id"`
		Subject        string `json:"subject"`
		Sender         uint64 `json:"sender"`
	}
	SuratMasukView struct {
		Classification sql.NullString `json:"classification"`
		CategoryID     sql.NullInt64  `json:"category_id"`
		PriorityID     sql.NullInt64  `json:"priority_id"`
		TypeID         sql.NullInt64  `json:"type_id"`
		Subject        sql.NullString `json:"subject"`
		Sender         sql.NullInt64  `json:"sender"`
	}
)

func GetJWTId(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	return id
}
func GetJWTName(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["username"].(string)
	return id
}
