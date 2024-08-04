package block

type PinkStainedGlass struct {
}

func (b PinkStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:pink_stained_glass", BlockProperties{}
}

func (b PinkStainedGlass) New(props BlockProperties) Block {
	return PinkStainedGlass{}
}