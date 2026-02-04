package tagged

import (
	"context"

	"github.com/CassianoDev/Xray-core5/common/net"
	"github.com/CassianoDev/Xray-core5/features/routing"
)

type DialFunc func(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
