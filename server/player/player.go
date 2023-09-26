package player

type Player struct {
	isHardCore       bool
	gameMode         byte
	previousGameMode int8

	viewDistance       int32
	simulationDistance int32

	Operator bool

	X, Y, Z    float64
	Yaw, Pitch float32
	OnGround   bool
}

func New() *Player {
	return &Player{}
}

func (p *Player) IsHardcore() bool {
	return p.isHardCore
}

func (p *Player) GameMode() byte {
	return p.gameMode
}

func (p *Player) PreviousGameMode() int8 {
	return p.previousGameMode
}

func (p *Player) ViewDistance() int32 {
	return p.viewDistance
}

func (p *Player) SimulationDistance() int32 {
	return p.viewDistance
}
