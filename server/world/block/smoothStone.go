package block

type SmoothStone struct {
}

func (b SmoothStone) Encode() (string, BlockProperties) {
	return "minecraft:smooth_stone", BlockProperties{}
}

func (b SmoothStone) New(props BlockProperties) Block {
	return SmoothStone{}
}