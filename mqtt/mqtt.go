package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/labstack/echo/v4"
	"os"
)

func InitMqtt(logger echo.Logger) { //nolint:typecheck
	var brokerUrl = os.Getenv("MQTT_URL")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerUrl)
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	logger.Info(fmt.Printf("Connecting to %s", brokerUrl))
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
