package handler

import (
	"aether/net"
	"aether/net/packet"
	"aether/net/packet/configuration"
	"aether/net/packet/play"
	"aether/server/session"
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
