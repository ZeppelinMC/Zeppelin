package block

type GreenWool struct {
}

func (b GreenWool) Encode() (string, BlockProperties) {
	return "minecraft:green_wool", BlockProperties{}
}

func (b GreenWool) New(props BlockProperties) Block {
	return GreenWool{}
}