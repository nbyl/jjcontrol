package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveAndLoadLightState(t *testing.T) {
	store := New()
	store.UpdateLightState(ON)
	assert.Equal(t, store.LightState, ON)
}
