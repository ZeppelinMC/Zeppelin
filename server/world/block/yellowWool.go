package block

type YellowWool struct {
}

func (b YellowWool) Encode() (string, BlockProperties) {
	return "minecraft:yellow_wool", BlockProperties{}
}

func (b YellowWool) New(props BlockProperties) Block {
	return YellowWool{}
}