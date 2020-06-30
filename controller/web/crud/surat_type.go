package crud

import (
	"PNM/model"
	"github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetSuratTypeALL(c echo.Context) (err error) {
	request := new(model.DefaultShowData)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	var count int

	if err := squirrel.Select("count(*)").From("mst_surat_type").RunWith(con).QueryRow().Scan(&count); err != nil {
		return err
	}

	rows, err := model.GetSuratType().
		Limit(request.Limit).
		Offset(request.Offset).
		RunWith(con).
		Query()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.CustomResponse{
		"data":  model.MapSuratTypeRows(rows),
		"count": count,
	})
}

func GetSuratTypeDropdown(c echo.Context) (err error) {
	con := model.InitDB()
	defer con.Close()

	rows, err := model.GetSuratType().
		Distinct().
		RunWith(con).
		Query()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.MapSuratTypeRows(rows))
}

func GetSuratTypeById(c echo.Context) (err error) {
	request := new(model.FindId)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	row := model.GetSuratType().Where(squirrel.Eq{"type_id": request.Id}).
		RunWith(con).
		QueryRow()

	return c.JSON(http.StatusOK, model.MapSingleSuratTypeRows(row))
}

func AddSuratType(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	if _, err := squirrel.Insert("mst_surat_type").
		Columns("type_name", "created_by", "created_at").
		Values(request.Name, id, time.Now()).RunWith(con).Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{Message: "Created blok"})
}

func EditSuratType(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	if _, err := squirrel.Update("mst_surat_type").
		Set("type_name", request.Name).
		Set("updated_by", id).
		Set("updated_at", time.Now()).
		Where("type_id = ?", request.Id).
		RunWith(con).
		Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{Message: "Updated blok"})
}

func DeleteSuratType(c echo.Context) (err error) {
	request := new(model.DefaultDelete)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	if _, err := squirrel.
		Delete("mst_surat_type").
		Where("type_id = ?", request.Id).
		RunWith(con).
		Exec();
		err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{Message: "Deleted blok"})
}
