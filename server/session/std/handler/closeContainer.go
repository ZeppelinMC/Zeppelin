package handler

import (
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdCloseContainer, handleCloseContainer)
}

func handleCloseContainer(s *std.StandardSession, p packet.Decodeable) {
	pk, ok := p.(*play.CloseContainer)
	if !ok {
		return
	}
	if s.WindowView.Load() == int32(pk.WindowId) {
		pos, w, ok := s.Dimension().WindowManager.Get(int32(pk.WindowId))
		s.WindowView.Store(0)
		if !ok {
			return
		}
		oldViewers := w.Viewers
		w.Viewers--

		x, y, z := pos[0], pos[1], pos[2]

		block, err := s.Dimension().Block(x, y, z)
		if err != nil {
			return
		}
		blockType, _ := block.Encode()

		if blockType == "minecraft:chest" {
			s.Broadcast().BlockAction(x, y, z, s.Player().Dimension(), 1, w.Viewers)
			if oldViewers > 0 && w.Viewers == 0 {
				// chest was closed by all players
				s.Broadcast().PlaySound(session.SoundEffect(
					"minecraft:block.chest.close", false, nil, play.SoundCategoryBlock, x, y, z, 1, 1,
				), s.Dimension().Name())
			}
		}
	}
}
