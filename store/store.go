package store

import (
	"encoding/json"
)

type PowerState int

const (
	ON      PowerState = 1
	OFF     PowerState = 2
	UNKNOWN PowerState = 3
)

type Store struct {
	LightState PowerState
}

func (s *Store) UpdateLightState(state PowerState) {
	s.LightState = state
}

func (s *Store) ToJson() (string, error) {
	value, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

func New() Store {
	return Store{
		LightState: UNKNOWN,
	}
}
