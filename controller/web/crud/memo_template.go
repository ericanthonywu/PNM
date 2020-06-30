package crud

import (
	"PNM/model"
	"fmt"
	"github.com/Masterminds/squirrel"
	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetMemoTemplateAll(c echo.Context) (err error) {
	request := new(model.DefaultShowData)
	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	var count int

	if err := squirrel.
		Select("count(*)").
		From("mst_memo_template").
		RunWith(con).
		QueryRow().
		Scan(&count); err != nil {
		return err
	}

	rows, err := model.GetSuratMemoTemplate().Limit(request.Limit).Offset(request.Offset).RunWith(con).Query()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.CustomResponse{
		"data":  model.MapMemoTemplateRows(rows),
		"count": count,
	})
}

func GetMemoTemplateById(c echo.Context) (err error) {
	request := new(model.FindId)

	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	row := model.GetSuratMemoTemplate().Where(squirrel.Eq{"template_id": request.Id}).RunWith(con).QueryRow()

	return c.JSON(http.StatusOK, model.MapSingleMemoTemplateRows(row))
}

func CreateMemoTemplate(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}
	fmt.Println(request)

	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// id := claims["id"].(string)

	if _, err := squirrel.Insert("mst_memo_template").
		Columns("memo_name", "memo_header", "memo_footer", "created_by", "created_at").
		Values(request.Name, request.Header, request.Footer, "superadmin", time.Now()).RunWith(con).Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{Message: "Created"})
}

func UpdateMemoTemplate(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// id := claims["id"].(string)

	if _, err := squirrel.Update("mst_memo_template").
		Set("memo_name", request.Name).
		Set("memo_header", request.Header).
		Set("memo_footer", request.Footer).
		// Set("updated_by", id).
		// Set("updated_at", time.Now()).
		RunWith(con).
		Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, model.Response{Message: "Updated"})
}

func DeleteMemoTemplate(c echo.Context) (err error) {
	request := new(model.DefaultDelete)

	con := model.InitDB()
	defer con.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	if _, err := squirrel.Delete("mst_memo_template").
		Where(squirrel.Eq{"template_id": request.Id}).
		RunWith(con).
		Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, model.Response{Message: "Deleted"})
}

func GetMemoTemplateDropdown(c echo.Context) (err error) {
	con := model.InitDB()
	defer con.Close()

	rows, err := model.GetSuratMemoTemplate().Distinct().RunWith(con).Query()

	return c.JSON(http.StatusOK, model.MapMemoTemplateRows(rows))
}
