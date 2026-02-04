package tcp

import (
	"github.com/CassianoDev/Xray-core5/common"
	"github.com/CassianoDev/Xray-core5/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
