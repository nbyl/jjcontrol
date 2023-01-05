package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nbyl/jjcontrol/backend/api"
	"github.com/nbyl/jjcontrol/backend/smarthome"
	"github.com/nbyl/jjcontrol/frontend"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v3"
	"os"
)

func main() {
	//nolint:typecheck
	smarthomeClient, err := smarthome.NewSmarthomeClient()
	if err != nil {
		log.Panic().Msg("Cannot connect to mqtt broker")
	}
	roomService := smarthome.NewRoomService(smarthomeClient)

	e := echo.New()
	e.Logger = lecho.New(os.Stdout)
	frontend.RegisterHandlers(e)
	api.NewRoomController(e, roomService)
	e.Logger.Fatal(e.Start(":8080"))
}
