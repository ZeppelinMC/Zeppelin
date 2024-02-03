package network

import "github.com/aimjel/minecraft/packet"

func HandleSwingArm(s *Session, p packet.Packet) {
	switch pk := p.(type) {
	case *packet.SwingArmServer:
		var anim animation
		switch pk.Hand {
		case 0: // main hand:
			anim = SwingMainArm
		case 1: // offhand:
			anim = SwingOffhand
		}
		s.b.BroadcastAnimation(s.eid, anim)
	}
}
