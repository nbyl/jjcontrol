package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/nbyl/jjcontrol/web"
	"net/http"
)

func main() {
	//nolint:typecheck
	e := echo.New()
	web.RegisterHandlers(e)
	//nolint:typecheck
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
