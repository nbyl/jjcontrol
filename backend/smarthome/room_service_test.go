package smarthome

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type SmarthomeClientMock struct {
	mock.Mock
}

func (s SmarthomeClientMock) RegisterListener(listener SmarthomeListener) { //nolint:govet
	s.Called(listener)
}

func (s SmarthomeClientMock) SendLightCommand(state PowerState) error { //nolint:govet
	s.Called(state)
	return nil
}

func TestSaveAndLoadLightState(t *testing.T) {
	clientMock := new(SmarthomeClientMock)
	clientMock.On("RegisterListener", mock.Anything).Return()

	service := NewRoomService(clientMock)
	service.UpdateLightState(ON)
	assert.Equal(t, service.GetLightState(), ON)
}

func TestSwitchLight(t *testing.T) {
	clientMock := new(SmarthomeClientMock)
	clientMock.On("RegisterListener", mock.Anything).Return()
	clientMock.On("SendLightCommand", ON).Return()

	service := NewRoomService(clientMock)
	err := service.SwitchLight(ON)
	assert.NoError(t, err)

	clientMock.AssertExpectations(t)
}
