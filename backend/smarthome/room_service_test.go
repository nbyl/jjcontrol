package smarthome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveAndLoadLightState(t *testing.T) {
	service := NewRoomService(&SmarthomeClient{listeners: []SmarthomeListener{}})
	service.UpdateLightState(ON)
	assert.Equal(t, service.LightState, ON)
}
