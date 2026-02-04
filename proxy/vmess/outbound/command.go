package outbound

import (

	"github.com/CassianoDev/Xray-core5/common/net"
	"github.com/CassianoDev/Xray-core5/common/protocol"
)

// As a stub command consumer.
func (h *Handler) handleCommand(dest net.Destination, cmd protocol.ResponseCommand) {
	switch cmd.(type) {
	default:
	}
}
