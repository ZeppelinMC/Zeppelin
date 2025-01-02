package block

type PinkWool struct {
}

func (b PinkWool) Encode() (string, BlockProperties) {
	return "minecraft:pink_wool", BlockProperties{}
}

func (b PinkWool) New(props BlockProperties) Block {
	return PinkWool{}
}