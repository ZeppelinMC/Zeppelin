package block

type SmoothSandstone struct {
}

func (b SmoothSandstone) Encode() (string, BlockProperties) {
	return "minecraft:smooth_sandstone", BlockProperties{}
}

func (b SmoothSandstone) New(props BlockProperties) Block {
	return SmoothSandstone{}
}