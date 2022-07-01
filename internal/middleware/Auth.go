package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if (len(c.Request().Header["Api-Key"]) > 0) {
			if (c.Request().Header["Api-Key"][0] == os.Getenv("API_KEY")) {
				return next(c)
			}
		}
		return c.JSON(http.StatusForbidden, "You are not authorized!")
	}
}