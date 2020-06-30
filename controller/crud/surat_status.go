package crud

import (
	"PNM/model"
	"github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetStatusByCode(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	rows := model.GetSuratStatus().
		Where(squirrel.Eq{"status_code": request.Code}).
		RunWith(con).
		QueryRow()

	return c.JSON(http.StatusOK, model.MapSingleSuratStatusRows(rows))
}

func EditSuratStatus(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)
	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	if _, err := squirrel.Update("mst_surat_status").
		Set("status_alias", request.Alias).
		Set("updated_by", id).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"status_id": request.Id}).
		RunWith(con).
		Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "Updated Blok",
	})
}

func GetAllSuratStatus(c echo.Context) (err error) {
	request := new(model.DefaultShowData)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	var count int

	if err := squirrel.Select("count(*)").From("mst_surat_status").RunWith(con).QueryRow().Scan(&count); err != nil {
		return err
	}

	rows, err := model.GetSuratStatus().
		Limit(request.Limit).
		Offset(request.Offset).
		RunWith(con).
		Query()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.CustomResponse{
		"data":  model.MapSuratStatusRows(rows),
		"count": count,
	})
}
