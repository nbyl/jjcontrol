package api

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/nbyl/jjcontrol/mqtt"
	"gitlab.com/nbyl/jjcontrol/store"
	"gitlab.com/nbyl/jjcontrol/web"
	"net/http"
)

var localState *store.Store

type Room struct {
	LightOn bool `json:"lightOn"`
}

func InitRestApi(e *echo.Echo, state *store.Store) { //nolint:typecheck
	localState = state
	web.RegisterHandlers(e)
	//nolint:typecheck
	e.GET("/api/room", func(c echo.Context) error {
		room := Room{
			LightOn: localState.LightState == store.ON,
		}

		return c.JSON(http.StatusOK, room)
	})
	e.PUT("/api/room", func(c echo.Context) error { // nolint:typecheck
		var room Room
		if err := c.Bind(&room); err != nil {
			e.Logger.Error(err)
			return c.String(http.StatusBadRequest, "bad request")
		}

		var lightState store.PowerState = store.OFF
		if room.LightOn {
			lightState = store.ON
		}
		mqtt.SendLightCommand(lightState)

		return c.String(http.StatusNoContent, "")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
