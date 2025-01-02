package block

type LightBlueStainedGlass struct {
}

func (b LightBlueStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_stained_glass", BlockProperties{}
}

func (b LightBlueStainedGlass) New(props BlockProperties) Block {
	return LightBlueStainedGlass{}
}