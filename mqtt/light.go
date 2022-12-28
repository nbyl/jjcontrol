package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gitlab.com/nbyl/jjcontrol/store"
	"os"
)

func subscribeToLightTopic(client mqtt.Client) {
	lightId := os.Getenv("TASMOTA_LIGHT_ID")

	var commandTopic = fmt.Sprintf("cmnd/%s/POWER", lightId)
	var statTopic = fmt.Sprintf("stat/%s/POWER", lightId)

	client.Subscribe(statTopic, 0, func(client mqtt.Client, message mqtt.Message) {
		var value = string(message.Payload())

		if value == "ON" {
			localState.UpdateLightState(store.ON)
		} else {
			localState.UpdateLightState(store.OFF)
		}
	})

	client.Publish(commandTopic, 0, false, "")
}
