package block

type SoulFire struct {
}

func (b SoulFire) Encode() (string, BlockProperties) {
	return "minecraft:soul_fire", BlockProperties{}
}

func (b SoulFire) New(props BlockProperties) Block {
	return SoulFire{}
}