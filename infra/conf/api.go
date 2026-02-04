package conf

import (
	"strings"

	"github.com/CassianoDev/Xray-core5/app/commander"
	loggerservice "github.com/CassianoDev/Xray-core5/app/log/command"
	observatoryservice "github.com/CassianoDev/Xray-core5/app/observatory/command"
	handlerservice "github.com/CassianoDev/Xray-core5/app/proxyman/command"
	routerservice "github.com/CassianoDev/Xray-core5/app/router/command"
	statsservice "github.com/CassianoDev/Xray-core5/app/stats/command"
	"github.com/CassianoDev/Xray-core5/common/errors"
	"github.com/CassianoDev/Xray-core5/common/serial"
)

type APIConfig struct {
	Tag      string   `json:"tag"`
	Listen   string   `json:"listen"`
	Services []string `json:"services"`
}

func (c *APIConfig) Build() (*commander.Config, error) {
	if c.Tag == "" {
		return nil, errors.New("API tag can't be empty.")
	}

	services := make([]*serial.TypedMessage, 0, 16)
	for _, s := range c.Services {
		switch strings.ToLower(s) {
		case "reflectionservice":
			services = append(services, serial.ToTypedMessage(&commander.ReflectionConfig{}))
		case "handlerservice":
			services = append(services, serial.ToTypedMessage(&handlerservice.Config{}))
		case "loggerservice":
			services = append(services, serial.ToTypedMessage(&loggerservice.Config{}))
		case "statsservice":
			services = append(services, serial.ToTypedMessage(&statsservice.Config{}))
		case "observatoryservice":
			services = append(services, serial.ToTypedMessage(&observatoryservice.Config{}))
		case "routingservice":
			services = append(services, serial.ToTypedMessage(&routerservice.Config{}))
		}
	}

	return &commander.Config{
		Tag:     c.Tag,
		Listen:  c.Listen,
		Service: services,
	}, nil
}
