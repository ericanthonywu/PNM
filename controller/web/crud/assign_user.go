package crud

import (
	"PNM/model"
	"github.com/Masterminds/squirrel"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddRecipient(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	for _, element := range request.UsersId {
		var username string
		if err := squirrel.Select("username").
			From("mst_user").
			Where(squirrel.Eq{"user_id": element}).
			RunWith(con).
			QueryRow().
			Scan(&username); err != nil {
			return err
		}
		if _, err := squirrel.Insert("rel_surat_recipient").
			Columns("surat_id", "recipient_name", "sign_status").
			Values(request.SuratId, username, 1).
			RunWith(con).
			Exec(); err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, model.Response{Message: "Inserted"})
}
