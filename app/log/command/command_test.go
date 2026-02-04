package command_test

import (
	"context"
	"testing"

	"github.com/CassianoDev/Xray-core5/app/dispatcher"
	"github.com/CassianoDev/Xray-core5/app/log"
	. "github.com/CassianoDev/Xray-core5/app/log/command"
	"github.com/CassianoDev/Xray-core5/app/proxyman"
	_ "github.com/CassianoDev/Xray-core5/app/proxyman/inbound"
	_ "github.com/CassianoDev/Xray-core5/app/proxyman/outbound"
	"github.com/CassianoDev/Xray-core5/common"
	"github.com/CassianoDev/Xray-core5/common/serial"
	"github.com/CassianoDev/Xray-core5/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
