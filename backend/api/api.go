package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nbyl/jjcontrol/backend/smarthome"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

type Room struct {
	Name    string `json:"name"`
	LightOn bool   `json:"lightOn"`
}

type RoomController struct {
	roomService smarthome.RoomService
}

func (r RoomController) GetState(c echo.Context) error {
	log.Info().Msgf("api:%p", r.roomService)
	room := Room{
		Name:    os.Getenv("ROOM_NAME"),
		LightOn: r.roomService.GetLightState() == smarthome.ON,
	}

	return c.JSON(http.StatusOK, room)
}

func (r RoomController) UpdateState(c echo.Context) error {
	var room Room
	if err := c.Bind(&room); err != nil {
		log.Err(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	//var lightState smarthome.PowerState = smarthome.OFF
	//if room.LightOn {
	//	lightState = smarthome.ON
	//}
	//localState.SendLightCommand(lightState)

	return c.String(http.StatusNoContent, "")
}

func NewRoomController(e *echo.Echo, roomService smarthome.RoomService) *RoomController {
	controller := RoomController{
		roomService: roomService,
	}

	e.GET("/api/room", controller.GetState)
	e.PUT("/api/room", controller.UpdateState)

	return &controller
}
