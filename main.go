package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nbyl/jjcontrol/backend/api"
	"github.com/nbyl/jjcontrol/backend/mqtt"
	"github.com/nbyl/jjcontrol/backend/store"
)

func main() {
	//nolint:typecheck
	e := echo.New()
	state := store.New()
	mqtt.InitMqtt(e.Logger, &state)
	api.InitRestApi(e, &state)
}
