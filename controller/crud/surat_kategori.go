package crud

import (
	"PNM/model"
	"github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetAllCategory(c echo.Context) (err error) {
	request := new(model.DefaultShowData)

	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	var count int

	if err := squirrel.
		Select("count(*)").
		From("mst_surat_category").
		RunWith(con).
		QueryRow().
		Scan(&count); err != nil {
		return err
	}

	rows, err := model.GetSuratCategory().
		Limit(request.Limit).
		Offset(request.Offset).
		RunWith(con).
		Query()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.CustomResponse{
		"data": model.MapSuratCategoryRows(rows),
		"count": count,
	})
}

func GetCategoryById(c echo.Context) (err error) {
	request := new(model.FindId)
	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	row := model.GetSuratCategory().
		Where("category_id = ?", request.Id).
		RunWith(con).
		QueryRow()

	return c.JSON(http.StatusOK, model.MapSingleSuratCategoryRows(row))
}

func AddCategory(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	if _, err := squirrel.Insert("mst_surat_category").
		Columns("category_name", "created_by", "created_at").
		Values(request.Name, id, time.Now()).
		RunWith(con).
		Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{
		Message: "Created blok",
	})
}

func EditCategory(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	if _, err := squirrel.Update("mst_surat_category").
		Set("category_name", request.Name).
		Set("updated_by", id).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"category_id": request.Id}).
		RunWith(con).
		Exec(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, model.Response{
		Message: "Updated blok",
	})
}

func DeleteCategory(c echo.Context) (err error) {
	request := new(model.DefaultDelete)

	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	if _, err := squirrel.
		Delete("mst_surat_category").
		Where(squirrel.Eq{"category_id": request.Id}).
		RunWith(con).
		Exec(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, model.Response{
		Message: "Deleted blok",
	})
}

func CategoryDropdown(c echo.Context) (err error) {
	con := model.InitDB()
	defer con.Close()

	rows, err := model.GetSuratCategory().Distinct().RunWith(con).Query()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.MapSuratCategoryRows(rows))
}
