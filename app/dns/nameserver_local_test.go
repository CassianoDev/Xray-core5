package dns_test

import (
	"context"
	"testing"
	"time"

	. "github.com/CassianoDev/Xray-core5/app/dns"
	"github.com/CassianoDev/Xray-core5/common"
	"github.com/CassianoDev/Xray-core5/features/dns"
)

func TestLocalNameServer(t *testing.T) {
	s := NewLocalNameServer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	ips, _, err := s.QueryIP(ctx, "google.com", dns.IPOption{
		IPv4Enable: true,
		IPv6Enable: true,
		FakeEnable: false,
	})
	cancel()
	common.Must(err)
	if len(ips) == 0 {
		t.Error("expect some ips, but got 0")
	}
}
