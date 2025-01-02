package block

type GrayStainedGlass struct {
}

func (b GrayStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:gray_stained_glass", BlockProperties{}
}

func (b GrayStainedGlass) New(props BlockProperties) Block {
	return GrayStainedGlass{}
}