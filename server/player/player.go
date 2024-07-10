package player

type Player struct {
	entityId int32
}

func NewPlayer(entityId int32) *Player {
	return &Player{entityId: entityId}
}

func (p *Player) EntityId() int32 {
	return p.entityId
}
