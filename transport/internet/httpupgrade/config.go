package httpupgrade

import (
	"github.com/CassianoDev/Xray-core5/common"
	"github.com/CassianoDev/Xray-core5/transport/internet"
)

func (c *Config) GetNormalizedPath() string {
	path := c.Path
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		return "/" + path
	}
	return path
}

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
