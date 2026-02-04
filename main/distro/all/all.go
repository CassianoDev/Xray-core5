package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/CassianoDev/Xray-core5/app/dispatcher"
	_ "github.com/CassianoDev/Xray-core5/app/proxyman/inbound"
	_ "github.com/CassianoDev/Xray-core5/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/CassianoDev/Xray-core5/app/commander"
	_ "github.com/CassianoDev/Xray-core5/app/log/command"
	_ "github.com/CassianoDev/Xray-core5/app/proxyman/command"
	_ "github.com/CassianoDev/Xray-core5/app/stats/command"

	// Developer preview services
	_ "github.com/CassianoDev/Xray-core5/app/observatory/command"

	// Other optional features.
	_ "github.com/CassianoDev/Xray-core5/app/dns"
	_ "github.com/CassianoDev/Xray-core5/app/dns/fakedns"
	_ "github.com/CassianoDev/Xray-core5/app/log"
	_ "github.com/CassianoDev/Xray-core5/app/metrics"
	_ "github.com/CassianoDev/Xray-core5/app/policy"
	_ "github.com/CassianoDev/Xray-core5/app/reverse"
	_ "github.com/CassianoDev/Xray-core5/app/router"
	_ "github.com/CassianoDev/Xray-core5/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/CassianoDev/Xray-core5/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/CassianoDev/Xray-core5/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/CassianoDev/Xray-core5/proxy/blackhole"
	_ "github.com/CassianoDev/Xray-core5/proxy/dns"
	_ "github.com/CassianoDev/Xray-core5/proxy/dokodemo"
	_ "github.com/CassianoDev/Xray-core5/proxy/freedom"
	_ "github.com/CassianoDev/Xray-core5/proxy/http"
	_ "github.com/CassianoDev/Xray-core5/proxy/loopback"
	_ "github.com/CassianoDev/Xray-core5/proxy/shadowsocks"
	_ "github.com/CassianoDev/Xray-core5/proxy/socks"
	_ "github.com/CassianoDev/Xray-core5/proxy/trojan"
	_ "github.com/CassianoDev/Xray-core5/proxy/vless/inbound"
	_ "github.com/CassianoDev/Xray-core5/proxy/vless/outbound"
	_ "github.com/CassianoDev/Xray-core5/proxy/vmess/inbound"
	_ "github.com/CassianoDev/Xray-core5/proxy/vmess/outbound"
	_ "github.com/CassianoDev/Xray-core5/proxy/wireguard"

	// Transports
	_ "github.com/CassianoDev/Xray-core5/transport/internet/grpc"
	_ "github.com/CassianoDev/Xray-core5/transport/internet/httpupgrade"
	_ "github.com/CassianoDev/Xray-core5/transport/internet/kcp"
	_ "github.com/CassianoDev/Xray-core5/transport/internet/reality"
	_ "github.com/CassianoDev/Xray-core5/transport/internet/splithttp"
	_ "github.com/CassianoDev/Xray-core5/transport/internet/tcp"
	_ "github.com/CassianoDev/Xray-core5/transport/internet/tls"
	_ "github.com/CassianoDev/Xray-core5/transport/internet/udp"
	_ "github.com/CassianoDev/Xray-core5/transport/internet/websocket"

	// Transport headers
	_ "github.com/CassianoDev/Xray-core5/transport/internet/headers/http"
	_ "github.com/CassianoDev/Xray-core5/transport/internet/headers/noop"

	// JSON & TOML & YAML
	_ "github.com/CassianoDev/Xray-core5/main/json"
	_ "github.com/CassianoDev/Xray-core5/main/toml"
	_ "github.com/CassianoDev/Xray-core5/main/yaml"

	// Load config from file or http(s)
	_ "github.com/CassianoDev/Xray-core5/main/confloader/external"

	// Commands
	_ "github.com/CassianoDev/Xray-core5/main/commands/all"
)
