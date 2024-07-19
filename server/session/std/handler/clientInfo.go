package handler

import (
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/metadata"
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
	var inf configuration.ClientInformation
	switch pk := p.(type) {
	case *configuration.ClientInformation:
		inf = *pk
		s.Player().SetClientInformation(*pk)
	case *play.ClientInformation:
		inf = pk.ClientInformation
		s.Player().SetClientInformation(pk.ClientInformation)
	default:
		return
	}

	new := metadata.Metadata{
		metadata.PlayerDisplayedSkinPartsIndex: metadata.Byte(inf.DisplayedSkinParts),
		metadata.PlayerMainHandIndex:           metadata.Byte(inf.MainHand),
	}
	s.Player().SetMetadataIndexes(new)
	s.Broadcast().EntityMetadata(s, new)
}
