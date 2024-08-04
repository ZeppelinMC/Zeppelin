package block

type SoulSand struct {
}

func (b SoulSand) Encode() (string, BlockProperties) {
	return "minecraft:soul_sand", BlockProperties{}
}

func (b SoulSand) New(props BlockProperties) Block {
	return SoulSand{}
}