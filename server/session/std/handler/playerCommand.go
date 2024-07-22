package handler

import (
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/metadata"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
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
				metadata.PoseIndex: metadata.Standing,
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
