package smarthome

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

type RoomService struct {
	client     *SmarthomeClient
	LightState PowerState
}

func (s *RoomService) UpdateLightState(state PowerState) {
	s.LightState = state
}

func (s *RoomService) ToJson() (string, error) {
	value, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

func NewRoomService(client *SmarthomeClient) *RoomService {
	service := RoomService{
		client:     client,
		LightState: UNKNOWN,
	}
	log.Info().Msgf("%p", &service)
	client.RegisterListener(&service)

	return &service
}
