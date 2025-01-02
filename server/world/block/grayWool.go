package block

type GrayWool struct {
}

func (b GrayWool) Encode() (string, BlockProperties) {
	return "minecraft:gray_wool", BlockProperties{}
}

func (b GrayWool) New(props BlockProperties) Block {
	return GrayWool{}
}