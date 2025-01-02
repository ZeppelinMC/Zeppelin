package block

type WhiteStainedGlass struct {
}

func (b WhiteStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:white_stained_glass", BlockProperties{}
}

func (b WhiteStainedGlass) New(props BlockProperties) Block {
	return WhiteStainedGlass{}
}