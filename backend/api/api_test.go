package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nbyl/jjcontrol/backend/smarthome"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type SmarthomeServiceMock struct {
	mock.Mock
}

func (s *SmarthomeServiceMock) SwitchLight(state smarthome.PowerState) error {
	args := s.Called(state)
	return args.Error(0)
}

func (s *SmarthomeServiceMock) GetLightState() smarthome.PowerState {
	args := s.Called()
	return args.Get(0).(smarthome.PowerState)
}

func (s *SmarthomeServiceMock) UpdateLightState(state smarthome.PowerState) {
}

func TestGetState(t *testing.T) {
	serviceMock := new(SmarthomeServiceMock)
	serviceMock.On("GetLightState").Return(smarthome.ON)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/room", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := NewRoomController(e, serviceMock)
	assert.NotNil(t, controller)

	if assert.NoError(t, controller.GetState(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"name\":\"\",\"lightOn\":true}\n", rec.Body.String())
	}
}
func TestUpdateState(t *testing.T) {
	serviceMock := new(SmarthomeServiceMock)
	serviceMock.On("SwitchLight", smarthome.ON).Return(nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/api/room", strings.NewReader("{\"lightOn\":true}\n"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := NewRoomController(e, serviceMock)
	assert.NotNil(t, controller)

	if assert.NoError(t, controller.UpdateState(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}

	serviceMock.AssertExpectations(t)
}
