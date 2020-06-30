package crud

import (
	"PNM/model"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetSuratPriorityALL(c echo.Context) (err error) {
	request := new(model.DefaultShowData)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	var count int

	if err := squirrel.
		Select("count(*)").
		From("mst_surat_priority").
		RunWith(con).
		QueryRow().
		Scan(&count); err != nil {
		return err
	}

	rows, err := model.GetSuratPrior().
		Limit(request.Limit).
		Offset(request.Offset).
		RunWith(con).
		Query()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.CustomResponse{
		"data":  model.MapSuratPriorRows(rows),
		"count": count,
	})
}

func GetSuratPriorityDropdown(c echo.Context) (err error) {
	con := model.InitDB()
	defer con.Close()

	rows, err := model.GetSuratPrior().
		Distinct().
		RunWith(con).
		Query()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.MapSuratPriorRows(rows))
}

func GetSuratPriorById(c echo.Context) (err error) {
	request := new(model.FindId)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	row := model.GetSuratPrior().Where(squirrel.Eq{"priority_id": request.Id}).RunWith(con).QueryRow()

	return c.JSON(http.StatusOK, model.MapSingleSuratPriorRows(row))
}

func AddSuratPriority(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	if err := c.Bind(request); err != nil {
		return err
	}

	if request.Name == "" {
		return echo.ErrBadRequest
	}

	con := model.InitDB()
	defer con.Close()

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	if _, err := squirrel.Insert("mst_surat_priority").
		Columns("priority_name", "created_by", "created_at").
		Values(request.Name, id, time.Now()).RunWith(con).Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{Message: "Created blok"})
}

func EditSuratPriority(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	fmt.Println(request)

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	if _, err := squirrel.Update("mst_surat_priority").
		Set("priority_name", request.Name).
		Set("updated_by", id).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"priority_id": request.Id}).
		RunWith(con).
		Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{Message: "Updated blok"})
}

func DeleteSuratPriority(c echo.Context) (err error) {
	request := new(model.DefaultDelete)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	if _, err := squirrel.
		Delete("mst_surat_priority").
		Where(squirrel.Eq{"priority_id": request.Id}).
		RunWith(con).
		Exec();
		err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, model.Response{Message: "Deleted blok"})
}
