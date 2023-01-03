package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nbyl/jjcontrol/api"
	"github.com/nbyl/jjcontrol/mqtt"
	"github.com/nbyl/jjcontrol/store"
)

func main() {
	//nolint:typecheck
	e := echo.New()
	state := store.New()
	mqtt.InitMqtt(e.Logger, &state)
	api.InitRestApi(e, &state)
}
