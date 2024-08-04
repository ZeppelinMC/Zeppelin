package block

type YellowStainedGlass struct {
}

func (b YellowStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:yellow_stained_glass", BlockProperties{}
}

func (b YellowStainedGlass) New(props BlockProperties) Block {
	return YellowStainedGlass{}
}