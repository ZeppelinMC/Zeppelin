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
		var newmd metadata.Metadata
		switch command.ActionId {
		case play.ActionIdStartSneaking:
			base := session.Player().MetadataIndex(metadata.BaseIndex).(metadata.Byte)
			base |= metadata.IsCrouching

			newmd = metadata.Metadata{
				metadata.BaseIndex: base,
				metadata.PoseIndex: metadata.Sneaking,
			}
		case play.ActionIdStopSneaking:
			base := session.Player().MetadataIndex(metadata.BaseIndex).(metadata.Byte)
			base &= ^metadata.IsCrouching

			newmd = metadata.Metadata{
				metadata.BaseIndex: base,
				metadata.PoseIndex: metadata.Sneaking,
			}
		case play.ActionIdStartSprinting:
			base := session.Player().MetadataIndex(metadata.BaseIndex).(metadata.Byte)
			base |= metadata.IsSprinting

			newmd = metadata.Metadata{
				metadata.BaseIndex: base,
			}
		case play.ActionIdStopSprinting:
			base := session.Player().MetadataIndex(metadata.BaseIndex).(metadata.Byte)
			base &= ^metadata.IsSprinting

			newmd = metadata.Metadata{
				metadata.BaseIndex: base,
			}
		}
		session.Player().SetMetadataIndexes(newmd)
		session.Broadcast().EntityMetadata(session, newmd)
	}
}
