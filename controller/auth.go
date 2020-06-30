package controller

import (
	"PNM/model"
	"github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

func Login(c echo.Context) (err error) {
	request := new(model.Login)
	user := new(model.User)

	if err := c.Bind(request); err != nil {
		return err
	}

	if request.Password == "" || request.Username == "" {
		return echo.ErrBadRequest
	}

	db := model.InitDB()

	if err := squirrel.Select("user_id", 
				"username", 
				"password",
				"email",
				"phone",
				"COALESCE(device,'')",
				"COALESCE(os,'')",
				"COALESCE(device_token,'')",
				"mst_user.position_id",
				"rep.position_name",
				"mst_user.division_id",
				"red.division_name",
		).
		From("mst_user").
		Where(squirrel.Eq{"username": request.Username}).
		LeftJoin("ref_employee_division red on mst_user.division_id = red.division_id").
		LeftJoin("ref_employee_position rep on mst_user.position_id = rep.position_id").
		RunWith(db).QueryRow().Scan(
			&user.Id, 
			&user.Username, 
			&user.Password,
			&user.Email,
			&user.Phone,
			&user.Device,
			&user.Os,
			&user.DeviceToken,
			&user.PositionId,
			&user.Position,
			&user.DivisionId,
			&user.Division); err != nil {
		return err
	}

	defer db.Close()

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password)); err != nil {
		return echo.ErrForbidden
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id"] = user.Id

	t, err := token.SignedString([]byte(os.Getenv("JWTTOKEN")))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.DefaultResponse{
		// Status: 1,
		Message: "Login Berhasil",
		Data: model.CustomResponse{
			"token": t,
			"user":  user,
		},
	})

}

func Migrate(c echo.Context) (err error) {
	db := model.InitDB()
	hashed, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("insert into mst_user (username,password,email,phone) values (?,?,?,?)", "superadmin", hashed, "ericanthonywu89@gmail.com", "081236391375")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{Message: "migrate successfully!"})
}
