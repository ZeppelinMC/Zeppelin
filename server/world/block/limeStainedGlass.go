package block

type LimeStainedGlass struct {
}

func (b LimeStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:lime_stained_glass", BlockProperties{}
}

func (b LimeStainedGlass) New(props BlockProperties) Block {
	return LimeStainedGlass{}
}