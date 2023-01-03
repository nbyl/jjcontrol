package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/nbyl/jjcontrol/backend/store"
	"os"
)

func subscribeToLightTopic(client mqtt.Client) {
	lightId := os.Getenv("TASMOTA_LIGHT_ID")

	var statTopic = fmt.Sprintf("stat/%s/POWER", lightId)

	client.Subscribe(statTopic, 0, func(client mqtt.Client, message mqtt.Message) {
		var value = string(message.Payload())

		if value == "ON" {
			localState.UpdateLightState(store.ON)
		} else {
			localState.UpdateLightState(store.OFF)
		}
	})

	client.Publish(getCommandTopic(), 0, false, "")
}

func getCommandTopic() string {
	lightId := os.Getenv("TASMOTA_LIGHT_ID")
	return fmt.Sprintf("cmnd/%s/POWER", lightId)
}

func SendLightCommand(state store.PowerState) {
	payload := "OFF"
	if state == store.ON {
		payload = "ON"
	}

	client.Publish(getCommandTopic(), 1, false, payload)
}
