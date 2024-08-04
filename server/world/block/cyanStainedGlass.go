package block

type CyanStainedGlass struct {
}

func (b CyanStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:cyan_stained_glass", BlockProperties{}
}

func (b CyanStainedGlass) New(props BlockProperties) Block {
	return CyanStainedGlass{}
}