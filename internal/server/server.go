package server

import (
	"net/http"

	"github.com/labstack/echo"
)

// MainPage : Hello
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "HelloWorld")
	}
}
