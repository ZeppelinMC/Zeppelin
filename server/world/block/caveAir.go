package block

type CaveAir struct {
}

func (b CaveAir) Encode() (string, BlockProperties) {
	return "minecraft:cave_air", BlockProperties{}
}

func (b CaveAir) New(props BlockProperties) Block {
	return CaveAir{}
}