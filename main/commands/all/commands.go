package all

import (
	"github.com/CassianoDev/Xray-core5/main/commands/all/api"
	"github.com/CassianoDev/Xray-core5/main/commands/all/convert"
	"github.com/CassianoDev/Xray-core5/main/commands/all/tls"
	"github.com/CassianoDev/Xray-core5/main/commands/base"
)

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		convert.CmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
		cmdMLDSA65,
		cmdMLKEM768,
		cmdVLESSEnc,
		cmdBuildMphCache,
	)
}
