package block

type SmoothQuartz struct {
}

func (b SmoothQuartz) Encode() (string, BlockProperties) {
	return "minecraft:smooth_quartz", BlockProperties{}
}

func (b SmoothQuartz) New(props BlockProperties) Block {
	return SmoothQuartz{}
}