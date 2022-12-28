package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/nbyl/jjcontrol/api"
	"gitlab.com/nbyl/jjcontrol/mqtt"
	"gitlab.com/nbyl/jjcontrol/store"
)

func main() {
	//nolint:typecheck
	e := echo.New()
	state := store.New()
	mqtt.InitMqtt(e.Logger, &state)
	api.InitRestApi(e, &state)
}
