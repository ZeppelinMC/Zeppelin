package block

type WhiteWool struct {
}

func (b WhiteWool) Encode() (string, BlockProperties) {
	return "minecraft:white_wool", BlockProperties{}
}

func (b WhiteWool) New(props BlockProperties) Block {
	return WhiteWool{}
}