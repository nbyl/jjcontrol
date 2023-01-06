package smarthome

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
	"os"
)

type SmarthomeClient struct {
	client    mqtt.Client
	listeners []SmarthomeListener
}

type SmarthomeListener interface {
	UpdateLightState(state PowerState)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Info().Msg("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Err(err).Msg("Connection lost")
}

func NewSmarthomeClient() (*SmarthomeClient, error) { //nolint:typecheck
	var brokerUrl = os.Getenv("MQTT_URL")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerUrl)
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	log.Info().Msgf("Connecting to %s", brokerUrl)

	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	smarthomeClient := &SmarthomeClient{
		client:    client,
		listeners: []SmarthomeListener{},
	}
	err := smarthomeClient.subscribeToLightTopic(client)
	if err != nil {
		return nil, err
	}

	return smarthomeClient, nil
}

func (s *SmarthomeClient) RegisterListener(listener SmarthomeListener) {
	s.listeners = append(s.listeners, listener)
}

func (s *SmarthomeClient) fireUpdateLightState(state PowerState) {
	log.Info().Msgf("%d", len(s.listeners))
	for _, listener := range s.listeners {
		log.Info().Msgf("fireUpdateLightState:%p", listener)
		listener.UpdateLightState(state)
	}
}

func (s *SmarthomeClient) subscribeToLightTopic(client mqtt.Client) error {
	statTopic := getStatTopic()
	token := client.Subscribe(statTopic, 0, func(client mqtt.Client, message mqtt.Message) {
		var value = string(message.Payload())

		log.Info().Msg(value)

		if value == "ON" {
			s.fireUpdateLightState(ON)
		} else {
			s.fireUpdateLightState(OFF)
		}
	})
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	client.Publish(getCommandTopic(), 0, false, "")
	return nil
}

func getCommandTopic() string {
	lightId := os.Getenv("TASMOTA_LIGHT_ID")
	return fmt.Sprintf("cmnd/%s/POWER", lightId)
}

func getStatTopic() string {
	lightId := os.Getenv("TASMOTA_LIGHT_ID")
	return fmt.Sprintf("stat/%s/POWER", lightId)
}

func (s *SmarthomeClient) SendLightCommand(state PowerState) {
	payload := "OFF"
	if state == ON {
		payload = "ON"
	}

	s.client.Publish(getCommandTopic(), 1, false, payload)
}
