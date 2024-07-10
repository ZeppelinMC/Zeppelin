package handler

import (
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/configuration"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/dynamitemc/aether/server/session"
)

func init() {
	session.Handlers[[2]int32{net.ConfigurationState, configuration.PacketIdClientInformation}] = handleClientInfo
	session.Handlers[[2]int32{net.PlayState, play.PacketIdClientInformation}] = handleClientInfo
}

func handleClientInfo(s *session.Session, p packet.Packet) {
	switch pk := p.(type) {
	case *configuration.ClientInformation:
		s.ClientInfo.Set(*pk)
	case *play.ClientInformation:
		s.ClientInfo.Set(pk.ClientInformation)
	}
}
