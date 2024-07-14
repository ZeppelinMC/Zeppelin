package handler

import (
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/configuration"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/dynamitemc/aether/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdClientInformation, handleClientInfo)
	std.RegisterHandler(net.ConfigurationState, configuration.PacketIdClientInformation, handleClientInfo)
}

func handleClientInfo(s *std.StandardSession, p packet.Packet) {
	switch pk := p.(type) {
	case *configuration.ClientInformation:
		s.Player().SetClientInformation(*pk)
	case *play.ClientInformation:
		s.Player().SetClientInformation(pk.ClientInformation)
	}
}
