package udp

import (
	"context"

	"github.com/CassianoDev/Xray-core5/common"
	"github.com/CassianoDev/Xray-core5/common/errors"
	"github.com/CassianoDev/Xray-core5/common/net"
	"github.com/CassianoDev/Xray-core5/transport/internet"
	"github.com/CassianoDev/Xray-core5/transport/internet/stat"
)

func init() {
	common.Must(internet.RegisterTransportDialer(protocolName,
		func(ctx context.Context, dest net.Destination, streamSettings *internet.MemoryStreamConfig) (stat.Connection, error) {
			var sockopt *internet.SocketConfig
			if streamSettings != nil {
				sockopt = streamSettings.SocketSettings
			}
			conn, err := internet.DialSystem(ctx, dest, sockopt)
			if err != nil {
				return nil, err
			}

			if streamSettings != nil && streamSettings.UdpmaskManager != nil {
				wrapper, ok := conn.(*internet.PacketConnWrapper)
				if !ok {
					conn.Close()
					return nil, errors.New("conn is not PacketConnWrapper")
				}

				raw := wrapper.Conn

				wrapper.Conn, err = streamSettings.UdpmaskManager.WrapPacketConnClient(raw)
				if err != nil {
					raw.Close()
					return nil, errors.New("mask err").Base(err)
				}
			}

			// TODO: handle dialer options
			return conn, nil
		}))
}
