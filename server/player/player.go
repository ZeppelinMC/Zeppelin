package player

import "github.com/dynamitemc/aether/atomic"

type Player struct {
	entityId int32

	x, y, z    atomic.AtomicValue[float64]
	yaw, pitch atomic.AtomicValue[float32]
}

func NewPlayer(entityId int32) *Player {
	return &Player{entityId: entityId}
}

func (p *Player) Position() (x, y, z float64) {
	return p.x.Get(), p.y.Get(), p.z.Get()
}

func (p *Player) Rotation() (yaw, pitch float32) {
	return p.yaw.Get(), p.pitch.Get()
}

func (p *Player) SetPosition(x, y, z float64) {
	p.x.Set(x)
	p.y.Set(y)
	p.z.Set(z)
}

func (p *Player) SetRotation(yaw, pitch float32) {
	p.yaw.Set(yaw)
	p.pitch.Set(pitch)
}

func (p *Player) EntityId() int32 {
	return p.entityId
}
