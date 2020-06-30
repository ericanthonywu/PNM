package crud

import (
	"PNM/model"
	"github.com/Masterminds/squirrel"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetDivisionDropdown(c echo.Context) (err error) {

	type ShowDivision struct {
		DivisionID   uint64 `json:"division_id"`
		DivisionCode string `json:"division_code"`
		DivisionName string `json:"division_name"`
	}

	var data []ShowDivision

	con := model.InitDB()
	defer con.Close()

	rows, err := squirrel.
		Select("division_id", "division_code", "division_name").
		From("ref_employee_division").
		Distinct().
		RunWith(con).Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tempData ShowDivision
		if err := rows.Scan(&tempData.DivisionID, &tempData.DivisionCode, &tempData.DivisionName); err != nil {
			return err
		}
		data = append(data, tempData)
	}
	return c.JSON(http.StatusOK, data)
}
