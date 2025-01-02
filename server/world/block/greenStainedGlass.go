package block

type GreenStainedGlass struct {
}

func (b GreenStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:green_stained_glass", BlockProperties{}
}

func (b GreenStainedGlass) New(props BlockProperties) Block {
	return GreenStainedGlass{}
}