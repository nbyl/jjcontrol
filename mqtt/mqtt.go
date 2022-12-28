package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/labstack/echo/v4"
	"gitlab.com/nbyl/jjcontrol/store"
	"os"
)

var localState *store.Store
var localLogger echo.Logger //nolint:typecheck
var client mqtt.Client

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	subscribeToLightTopic(client)
	localLogger.Info("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	localLogger.Info("Connect lost: ", err)
}

func InitMqtt(logger echo.Logger, state *store.Store) { //nolint:typecheck
	localLogger = logger
	localState = state

	var brokerUrl = os.Getenv("MQTT_URL")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerUrl)
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client = mqtt.NewClient(opts)
	logger.Info(fmt.Printf("Connecting to %s", brokerUrl))
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
