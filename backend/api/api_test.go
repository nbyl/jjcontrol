package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nbyl/jjcontrol/backend/smarthome"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SmarthomeControllerMock struct {
	mock.Mock
}

func (s *SmarthomeControllerMock) GetLightState() smarthome.PowerState {
	return smarthome.ON
}

func (s *SmarthomeControllerMock) UpdateLightState(state smarthome.PowerState) {
}

func TestGetState(t *testing.T) {
	controllerMock := new(SmarthomeControllerMock)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := NewRoomController(e, controllerMock)
	assert.NotNil(t, controller)

	if assert.NoError(t, controller.GetState(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"name\":\"\",\"lightOn\":true}\n", rec.Body.String())
	}
}
