package api

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/nbyl/jjcontrol/store"
	"gitlab.com/nbyl/jjcontrol/web"
	"net/http"
)

var localState *store.Store

func InitRestApi(e *echo.Echo, state *store.Store) { //nolint:typecheck
	localState = state
	web.RegisterHandlers(e)
	//nolint:typecheck
	e.GET("/api/state", func(c echo.Context) error {
		value, err := localState.ToJson()
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.String(http.StatusOK, value)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
