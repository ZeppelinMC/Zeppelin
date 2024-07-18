package handler

import (
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/metadata"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/dynamitemc/aether/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdPlayerCommand, handlePlayerCommand)
}

func handlePlayerCommand(session *std.StandardSession, pk packet.Packet) {
	if command, ok := pk.(*play.PlayerCommand); ok {
		switch command.ActionId {
		case play.ActionIdStartSneaking:
			session.Player().SetMetadataIndex(metadata.PoseIndex, metadata.Sneaking)
			base := session.Player().MetadataIndex(metadata.BaseIndex).(metadata.Byte)
			base |= metadata.IsCrouching
			session.Player().SetMetadataIndex(metadata.BaseIndex, base)

			session.Broadcast().EntityMetadata(session, session.Player().Metadata())
		case play.ActionIdStopSneaking:
			session.Player().SetMetadataIndex(metadata.PoseIndex, metadata.Standing)
			base := session.Player().MetadataIndex(metadata.BaseIndex).(metadata.Byte)
			base &= ^metadata.IsCrouching
			session.Player().SetMetadataIndex(metadata.BaseIndex, base)

			session.Broadcast().EntityMetadata(session, session.Player().Metadata())
		case play.ActionIdStartSprinting:
			base := session.Player().MetadataIndex(metadata.BaseIndex).(metadata.Byte)
			base |= metadata.IsSprinting
			session.Player().SetMetadataIndex(metadata.BaseIndex, base)
			session.Broadcast().EntityMetadata(session, session.Player().Metadata())
		case play.ActionIdStopSprinting:
			base := session.Player().MetadataIndex(metadata.BaseIndex).(metadata.Byte)
			base &= ^metadata.IsSprinting
			session.Player().SetMetadataIndex(metadata.BaseIndex, base)
			session.Broadcast().EntityMetadata(session, session.Player().Metadata())
		}
	}
}
