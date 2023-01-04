package smarthome

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
)

type mosquittoContainer struct { //nolint:unused
	container testcontainers.Container
	url       string
}

func setupMosquitto() (*mosquittoContainer, error) { //nolint:unused
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "eclipse-mosquitto:2.0.15",
		ExposedPorts: []string{"1883/tcp"},

		WaitingFor: wait.ForListeningPort("1883/tcp"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return nil, err
	}

	host, _ := container.Host(ctx)

	p, _ := container.MappedPort(ctx, "1883/tcp")
	port := p.Int()

	return &mosquittoContainer{
		container: container,
		url:       fmt.Sprintf("tcp://%s:%d", host, port),
	}, nil
}

func teardownMosquitto(t *testing.T, container *mosquittoContainer) { //nolint:unused
	if err := container.container.Terminate(context.Background()); err != nil {
		t.Fatalf("failed to terminate container: %s", err.Error())
	}
}

func TestNewSmarthomeClient(t *testing.T) {
	t.Skip()
	container, err := setupMosquitto()
	if err != nil {
		t.Error(err)
	}
	defer teardownMosquitto(t, container)

	os.Setenv("MQTT_URL", container.url)

	_, err = NewSmarthomeClient()
	assert.NoError(t, err)
}
