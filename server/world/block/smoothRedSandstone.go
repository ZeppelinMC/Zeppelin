package block

type SmoothRedSandstone struct {
}

func (b SmoothRedSandstone) Encode() (string, BlockProperties) {
	return "minecraft:smooth_red_sandstone", BlockProperties{}
}

func (b SmoothRedSandstone) New(props BlockProperties) Block {
	return SmoothRedSandstone{}
}