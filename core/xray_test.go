package core_test

import (
	"testing"

	"github.com/CassianoDev/Xray-core5/app/dispatcher"
	"github.com/CassianoDev/Xray-core5/app/proxyman"
	"github.com/CassianoDev/Xray-core5/common"
	"github.com/CassianoDev/Xray-core5/common/net"
	"github.com/CassianoDev/Xray-core5/common/protocol"
	"github.com/CassianoDev/Xray-core5/common/serial"
	"github.com/CassianoDev/Xray-core5/common/uuid"
	. "github.com/CassianoDev/Xray-core5/core"
	"github.com/CassianoDev/Xray-core5/features/dns"
	"github.com/CassianoDev/Xray-core5/features/dns/localdns"
	_ "github.com/CassianoDev/Xray-core5/main/distro/all"
	"github.com/CassianoDev/Xray-core5/proxy/dokodemo"
	"github.com/CassianoDev/Xray-core5/proxy/vmess"
	"github.com/CassianoDev/Xray-core5/proxy/vmess/outbound"
	"github.com/CassianoDev/Xray-core5/testing/servers/tcp"
	"google.golang.org/protobuf/proto"
)

func TestXrayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	}, false)
	instance.AddFeature(localdns.New())
	<-wait
}

func TestXrayClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{
						Range: []*net.PortRange{net.SinglePortRange(port)},
					},
					Listen: net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address:  net.NewIPOrDomain(net.LocalHostIP),
					Port:     uint32(0),
					Networks: []net.Network{net.Network_TCP},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: &protocol.ServerEndpoint{
						Address: net.NewIPOrDomain(net.LocalHostIP),
						Port:    uint32(0),
						User:  &protocol.User{
							Account: serial.ToTypedMessage(&vmess.Account{
								Id: userID.String(),
							}),
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
