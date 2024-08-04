package block

type BlackStainedGlass struct {
}

func (b BlackStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:black_stained_glass", BlockProperties{}
}

func (b BlackStainedGlass) New(props BlockProperties) Block {
	return BlackStainedGlass{}
}