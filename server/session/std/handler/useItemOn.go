package handler

import (
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	"github.com/zeppelinmc/zeppelin/server/world/block"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdUseItemOn, handleUseItemOn)
}

func handleUseItemOn(s *std.StandardSession, pk packet.Packet) {
	use, ok := pk.(*play.UseItemOn)
	if !ok {
		return
	}
	dimension := s.Dimension()
	b, err := dimension.Block(use.BlockX, use.BlockY, use.BlockZ)
	if err != nil {
		return
	}

	usable, ok := b.(block.Usable)
	if !ok {
		return
	}
	usable.Use(s, *use, dimension)
}
