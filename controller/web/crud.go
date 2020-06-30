package web

import (
	"PNM/model"
	"github.com/Masterminds/squirrel"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func ShowUser(c echo.Context) error {
	con := model.InitDB()
	defer con.Close()

	request := new(model.DefaultShowData)
	if err := c.Bind(request); err != nil {
		return err
	}

	var u []model.UserResult

	//rows, err := model.GetUser().
	//	RunWith(con).
	//	Query()

	//if err != nil {
	//	return err
	//}



	return c.JSON(http.StatusOK, u)
}

func InsertUser(c echo.Context) error {
	request := model.NewUser()
	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if _, err := squirrel.Insert("mst_user").
		Columns("username", "password", "email", "phone", "is_blocked", "position_id", "division_id").
		Values(request.Username, password, request.Email, request.Phone, request.IsBlocked, request.PositionId, request.DivisionId).
		RunWith(con).
		Exec();
	err != nil {
		return err
	}


	//if _, err := con.Exec("insert into mst_user "+
	//	"(username, password, email, phone, is_blocked, position_id, division_id) values "+
	//	"(?,?,?,?,?,?,?)",
	//	request.Username,
	//	password,
	//	request.Email,
	//	request.Phone,
	//	request.IsBlocked,
	//	request.PositionId,
	//	request.DivisionId); err != nil {
	//	return err
	//}

	return c.JSON(http.StatusCreated, model.Response{
		Message: "Created blok",
	})
}

func UpdateUser(c echo.Context) error {
	request := new(model.User)
	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	//if _, err := con.Exec("update mst_user set"+
	//	"username = ?"+
	//	"email = ?"+
	//	"phone = ?"+
	//	"is_blocked = ?"+
	//	"position_id = ?"+
	//	"division_id = ?"+
	//	"where user_id = ?",
	//	request.Username,
	//	request.Email,
	//	request.Phone,
	//	request.IsBlocked,
	//	request.PositionId,
	//	request.DivisionId,
	//	request.Id); err != nil {
	//	return err
	//}
	return c.JSON(http.StatusCreated, model.Response{
		Message: "Updated blok",
	})
}

func DeleteUser(c echo.Context) error {
	request := new(model.DefaultDelete)

	db := model.InitDB()
	defer db.Close()

	if err := c.Bind(request); err != nil {
		return err
	}

	if _, err := db.Exec("delete from mst_user where user_id = ?", request.Id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "udah ke delete blok",
	})
}
