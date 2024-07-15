package handler

import (
	"fmt"

	"github.com/dynamitemc/aether/net"
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
			session.Broadcast().EntityMetadata(session, map[byte]any{
				6: play.Sneaking,
			})
			fmt.Println(session.Username(), "started sneaking")
		case play.ActionIdStopSneaking:
			session.Broadcast().EntityMetadata(session, map[byte]any{
				6: play.Standing,
			})
			//TODO add entity data (i0), for crouching, sprinting etc
		}
	}
}
