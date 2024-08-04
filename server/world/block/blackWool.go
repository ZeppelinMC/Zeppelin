package block

type BlackWool struct {
}

func (b BlackWool) Encode() (string, BlockProperties) {
	return "minecraft:black_wool", BlockProperties{}
}

func (b BlackWool) New(props BlockProperties) Block {
	return BlackWool{}
}