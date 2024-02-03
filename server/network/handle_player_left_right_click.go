package network

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/world/entity"
)

func HandlePlayerLeftRightClick(s *Session, pk packet.Packet) {
	switch p := pk.(type) {

	case *packet.InteractServer:
		//todo check this entity id is within a 4 block radius
		mu.RLock()
		en := entities[p.EntityID]
		mu.RUnlock()

		switch p.Type {

		case 0, 2:
			s.state.Interact(en)

			//attack
		case 1:
			e, ok := reg.DamageType.Lookup("minecraft:player_attack")
			if !ok {
				return
			}

			d := &packet.DamageEvent{
				EntityID:       p.EntityID,
				SourceTypeID:   int32(e.ID),
				SourceCauseID:  s.eid + 1,
				SourceDirectID: s.eid + 1,
			}

			if a, ok := en.(entity.Attacker); ok {
				s.state.Attack(a)
			}

			s.conn.SendPacket(d)

		}
	}

}
