package block

type VoidAir struct {
}

func (b VoidAir) Encode() (string, BlockProperties) {
	return "minecraft:void_air", BlockProperties{}
}

func (b VoidAir) New(props BlockProperties) Block {
	return VoidAir{}
}