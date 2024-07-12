package handler

import (
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/configuration"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/dynamitemc/aether/server/session"
)

var (
	_ = session.RegisterHandler(net.ConfigurationState, configuration.PacketIdClientInformation, handleClientInfo)
	_ = session.RegisterHandler(net.PlayState, play.PacketIdClientInformation, handleClientInfo)
)

func handleClientInfo(s *session.Session, p packet.Packet) {
	switch pk := p.(type) {
	case *configuration.ClientInformation:
		s.ClientInfo.Set(*pk)
	case *play.ClientInformation:
		s.ClientInfo.Set(pk.ClientInformation)
	}
}
