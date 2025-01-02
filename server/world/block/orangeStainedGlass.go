package block

type OrangeStainedGlass struct {
}

func (b OrangeStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:orange_stained_glass", BlockProperties{}
}

func (b OrangeStainedGlass) New(props BlockProperties) Block {
	return OrangeStainedGlass{}
}