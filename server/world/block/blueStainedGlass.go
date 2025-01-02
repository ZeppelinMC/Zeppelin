package block

type BlueStainedGlass struct {
}

func (b BlueStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:blue_stained_glass", BlockProperties{}
}

func (b BlueStainedGlass) New(props BlockProperties) Block {
	return BlueStainedGlass{}
}