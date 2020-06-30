package middleware

import (
	"PNM/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PNMMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if true {
			return next(c)
		}

		return c.JSON(http.StatusBadRequest,model.ErrorResponse{Message: "bad request"})
	}
}
