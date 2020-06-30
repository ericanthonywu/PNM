package crud

import (
	"PNM/model"
	"github.com/Masterminds/squirrel"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetALLUser(c echo.Context) (err error) {
	request := new(model.DefaultShowData)
	con := model.InitDB()
	if err := c.Bind(request); err != nil {
		return err
	}
	defer con.Close()

	var count int

	if err := squirrel.Select("count(*) as total_rows").From("mst_user").
		RunWith(con).
		QueryRow().
		Scan(&count); err != nil {
		return err
	}

	rows, err := model.GetUser().
		Limit(request.Limit).
		Offset(request.Offset).
		RunWith(con).
		Query()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.CustomResponse{
		"data":  model.MapUserRows(rows),
		"total": count,
	})
}

func GetUserById(c echo.Context) (err error) {
	request := new(model.FindId)
	if err := c.Bind(request); err != nil {
		return err
	}
	con := model.InitDB()
	defer con.Close()

	rows := model.GetUser().Where(squirrel.Eq{"user_id": request.Id}).RunWith(con).QueryRow()

	return c.JSON(http.StatusOK, model.MapSingleUserRows(rows))
}

func EditUser(c echo.Context) (err error) {
	request := new(model.ToogleIsBlocked)
	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	if _, err := squirrel.Update("mst_user").
		Where(squirrel.Eq{"user_id": request.Id}).
		Set("is_blocked", request.IsBlocked).
		RunWith(con).
		Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "Updated blok",
	})
}

func GetListUser(c echo.Context) (err error) {
	type data struct {
		Username string `json:"username"`
		UserId   uint64 `json:"user_id"`
		Division string `json:"division"`
		Position string `json:"position"`
	}
	var listData []data

	con := model.InitDB()
	defer con.Close()

	rows, err := squirrel.Select(
		"mst_user.username",
		"mst_user.user_id",
		"ref_employee_division.division_name as division",
		"ref_employee_position.position_name as position",
	).
		LeftJoin("ref_employee_division on mst_user.division_id = ref_employee_division.division_id").
		LeftJoin("ref_employee_position on mst_user.position_id = ref_employee_position.position_id").
		From("mst_user").
		RunWith(con).
		Query()

	if err != nil {
		return err
	}

	for rows.Next() {
		var data data
		if err := rows.Scan(&data.Username,
			&data.UserId,
			&data.Division,
			&data.Position,
		); err != nil {
			return err
		}
		listData = append(listData, data)
	}

	return c.JSON(http.StatusOK, listData)
}
