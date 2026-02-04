package conf_test

import (
	"testing"

	"github.com/CassianoDev/Xray-core5/common/net"
	"github.com/CassianoDev/Xray-core5/common/protocol"
	. "github.com/CassianoDev/Xray-core5/infra/conf"
	"github.com/CassianoDev/Xray-core5/proxy/freedom"
	"github.com/CassianoDev/Xray-core5/transport/internet"
)

func TestFreedomConfig(t *testing.T) {
	creator := func() Buildable {
		return new(FreedomConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"domainStrategy": "AsIs",
				"redirect": "127.0.0.1:3366",
				"userLevel": 1
			}`,
			Parser: loadJSON(creator),
			Output: &freedom.Config{
				DomainStrategy: internet.DomainStrategy_AS_IS,
				DestinationOverride: &freedom.DestinationOverride{
					Server: &protocol.ServerEndpoint{
						Address: &net.IPOrDomain{
							Address: &net.IPOrDomain_Ip{
								Ip: []byte{127, 0, 0, 1},
							},
						},
						Port: 3366,
					},
				},
				UserLevel: 1,
			},
		},
	})
}
