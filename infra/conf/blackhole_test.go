package conf_test

import (
	"testing"

	"github.com/CassianoDev/Xray-core5/common/serial"
	. "github.com/CassianoDev/Xray-core5/infra/conf"
	"github.com/CassianoDev/Xray-core5/proxy/blackhole"
)

func TestHTTPResponseJSON(t *testing.T) {
	creator := func() Buildable {
		return new(BlackholeConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"response": {
					"type": "http"
				}
			}`,
			Parser: loadJSON(creator),
			Output: &blackhole.Config{
				Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
			},
		},
		{
			Input:  `{}`,
			Parser: loadJSON(creator),
			Output: &blackhole.Config{},
		},
	})
}
