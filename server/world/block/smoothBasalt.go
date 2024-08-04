package block

type SmoothBasalt struct {
}

func (b SmoothBasalt) Encode() (string, BlockProperties) {
	return "minecraft:smooth_basalt", BlockProperties{}
}

func (b SmoothBasalt) New(props BlockProperties) Block {
	return SmoothBasalt{}
}