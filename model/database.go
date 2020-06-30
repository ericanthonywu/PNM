package model

import (
	dto "PNM/database"
	"database/sql"
	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/random"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func InitDB() *sql.DB {
	db, err := sql.Open(os.Getenv("DBDRIVER"), os.Getenv("DBURL"))
	if err != nil {
		panic(err)
	}

	return db
}

func InsertImage(header *multipart.FileHeader, dest string) (string, error) {

	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	dest = basepath + "/" + dest

	src, err := header.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	randoms := random.New()

	filename := time.Now().Format(time.RFC3339Nano) + randoms.String(30, random.Alphanumeric) + header.Filename

	dst, err := os.Create(dest + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return filename, nil
}

type (
	User struct {
		Id               uint64    `json:"id"`
		Username         *string   `json:"username" validate:"required"`
		Password         []byte    `json:"-" validate:"required"`
		Email            *string   `json:"email"`
		Phone            *string   `json:"phone"`
		Device           string    `json:"device"`
		Os               string    `json:"os"`
		DeviceToken      string    `json:"device_token"`
		VerifKey         string    `json:"verif_key"`
		VerifExpiredDate time.Time `json:"verif_expired_date"`
		OtpNumber        string    `json:"otp_number"`
		Apikey           string    `json:"apikey"`
		LastLoginIp      string    `json:"last_login_ip"`
		IsBlocked        bool      `json:"is_blocked"`
		PositionId       uint64    `json:"position_id"`
		DivisionId       uint64    `json:"division_id"`
		CreatedBy        string    `json:"created_by"`
		CreatedAt        time.Time `json:"created_by"`
		UpdatedBy        string    `json:"updated_by"`
	}
	LogActivity struct {
		LogId    uint64 `json:"log_id"`
		UserId   uint64 `json:"user_id"`
		Username string `json:"username"`
	}

	UserResult struct {
		Id       *uint64 `json:"id"`
		Username *string `json:"username"`
		Email    *string `json:"email"`
		Phone    *string `json:"phone"`
		Position *string `json:"position"`
		Division *string `json:"division"`
	}
)

func NewUser() *User {
	return &User{CreatedAt: time.Now()}
}

func MapDefaultRows(rows *sql.Rows) []MasterCrudDefault {
	defer rows.Close()
	var u []MasterCrudDefault

	for rows.Next() {
		u = append(u, MapSingleDefaultRows(rows))
	}
	return u
}

func MapSingleDefaultRows(rows *sql.Rows) MasterCrudDefault {
	var temp MasterCrudDefault

	if err := rows.Scan(&temp.Id, &temp.Name); err != nil {
		panic(err)
	}

	return temp
}

func GetUser() squirrel.SelectBuilder {
	return squirrel.Select(
		"mst_user.user_id",
		"mst_user.username",
		"mst_user.email",
		"mst_user.phone",
		"ref_employee_position.position_name as position",
		"ref_employee_division.division_name as division").
		From("mst_user").
		LeftJoin("ref_employee_division on mst_user.division_id = ref_employee_division.division_id").
		LeftJoin("ref_employee_position on mst_user.position_id = ref_employee_position.position_id")
}

func MapUserRows(rows *sql.Rows) []UserResult {
	defer rows.Close()
	var u []UserResult

	for rows.Next() {
		u = append(u, MapSingleUserRows(rows))
	}
	return u
}

func MapSingleUserRows(rows squirrel.RowScanner) UserResult {
	var userTemp UserResult

	if err := rows.Scan(&userTemp.Id, &userTemp.Username, &userTemp.Email, &userTemp.Phone, &userTemp.Position, &userTemp.Division); err != nil {
		panic(err)
	}

	return userTemp
}

func GetSuratPrior() squirrel.SelectBuilder {
	return squirrel.
		Select("priority_id", "priority_name").
		From("mst_surat_priority")
}

func MapSuratPriorRows(rows *sql.Rows) []dto.MstSuratPriority {
	defer rows.Close()
	var u []dto.MstSuratPriority

	for rows.Next() {
		u = append(u, MapSingleSuratPriorRows(rows))
	}
	return u
}

func MapSingleSuratPriorRows(rows squirrel.RowScanner) dto.MstSuratPriority {
	var temp dto.MstSuratPriority

	if err := rows.Scan(&temp.PriorityID, &temp.PriorityName); err != nil {
		panic(err)
	}

	return temp
}

func GetSuratCategory() squirrel.SelectBuilder {
	return squirrel.
		Select("category_id", "category_name").
		From("mst_surat_category")
}

func MapSuratCategoryRows(rows *sql.Rows) []dto.MstSuratCategory {
	defer rows.Close()
	var u []dto.MstSuratCategory

	for rows.Next() {
		u = append(u, MapSingleSuratCategoryRows(rows))
	}
	return u
}

func MapSingleSuratCategoryRows(rows squirrel.RowScanner) dto.MstSuratCategory {
	var temp dto.MstSuratCategory

	if err := rows.Scan(&temp.CategoryID, &temp.CategoryName); err != nil {
		panic(err)
	}

	return temp
}

func GetSuratType() squirrel.SelectBuilder {
	return squirrel.
		Select("type_id", "type_name").
		From("mst_surat_type")
}

func MapSuratTypeRows(rows *sql.Rows) []dto.MstSuratType {
	defer rows.Close()
	var u []dto.MstSuratType

	for rows.Next() {
		u = append(u, MapSingleSuratTypeRows(rows))
	}
	return u
}

func MapSingleSuratTypeRows(rows squirrel.RowScanner) dto.MstSuratType {
	var temp dto.MstSuratType

	if err := rows.Scan(&temp.TypeID, &temp.TypeName); err != nil {
		panic(err)
	}

	return temp
}

func GetSuratStatus() squirrel.SelectBuilder {
	return squirrel.
		Select("status_id", "status_code", "status", "status_alias").
		From("mst_surat_status")
}

func MapSuratStatusRows(rows *sql.Rows) []dto.MstSuratStatus {
	defer rows.Close()
	var u []dto.MstSuratStatus

	for rows.Next() {
		u = append(u, MapSingleSuratStatusRows(rows))
	}
	return u
}

func MapSingleSuratStatusRows(rows squirrel.RowScanner) dto.MstSuratStatus {
	var temp dto.MstSuratStatus

	if err := rows.Scan(&temp.StatusID, &temp.StatusCode, &temp.Status, &temp.StatusAlias); err != nil {
		panic(err)
	}

	return temp
}

func GetSuratMemoTemplate() squirrel.SelectBuilder {
	return squirrel.
		Select("template_id", "memo_name", "memo_header", "memo_footer").
		From("mst_memo_template")
}

func MapMemoTemplateRows(rows *sql.Rows) []dto.MstMemoTemplate {
	defer rows.Close()
	var u []dto.MstMemoTemplate

	for rows.Next() {
		u = append(u, MapSingleMemoTemplateRows(rows))
	}
	return u
}

func MapSingleMemoTemplateRows(rows squirrel.RowScanner) dto.MstMemoTemplate {
	var temp dto.MstMemoTemplate

	if err := rows.Scan(&temp.TemplateID, &temp.MemoName, &temp.MemoHeader, &temp.MemoFooter); err != nil {
		panic(err)
	}

	return temp
}
