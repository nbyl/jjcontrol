package smarthome

type RoomService interface {
	GetLightState() PowerState
	SwitchLight(state PowerState) error
	UpdateLightState(state PowerState)
}

type RoomServiceImplementation struct {
	client     SmarthomeClient
	LightState PowerState
}

func (s *RoomServiceImplementation) SwitchLight(state PowerState) error {
	return s.client.SendLightCommand(state)
}

func (s *RoomServiceImplementation) GetLightState() PowerState {
	return s.LightState
}

func (s *RoomServiceImplementation) UpdateLightState(state PowerState) {
	s.LightState = state
}

func NewRoomService(client SmarthomeClient) RoomService {
	service := RoomServiceImplementation{
		client:     client,
		LightState: UNKNOWN,
	}
	client.RegisterListener(&service)

	return &service
}
