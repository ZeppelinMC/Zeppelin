package handler

import (
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdCloseContainer, handleCloseContainer)
}

func handleCloseContainer(s *std.StandardSession, p packet.Packet) {
	pk, ok := p.(*play.CloseContainer)
	if !ok {
		return
	}
	if s.WindowView.Get() == int32(pk.WindowId) {
		pos, w, ok := s.Dimension().WindowManager.Get(int32(pk.WindowId))
		s.WindowView.Set(0)
		if !ok {
			return
		}
		w.Viewers--
		if w.Viewers < 0 {
			w.Viewers = 0
		}
		s.Broadcast().BlockAction(pos[0], pos[1], pos[2], s.Player().Dimension(), 1, w.Viewers)
	}
}
