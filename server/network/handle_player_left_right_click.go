package network

import "github.com/aimjel/minecraft/packet"

func HandlePlayerLeftRightClick(s *Session, pk packet.Packet) {
	switch p := pk.(type) {

	case *packet.InteractServer:
		mu.RLock()
		en := entities[p.EntityID]
		mu.RUnlock()

		switch p.Type {

		case 0, 2:
			//todo check within a 4unit radius without visible obstruction
			s.state.Interact(en)

			//attack
		case 1:

		}
	}

}
