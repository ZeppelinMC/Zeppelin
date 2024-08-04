package block

type CutSandstone struct {
}

func (b CutSandstone) Encode() (string, BlockProperties) {
	return "minecraft:cut_sandstone", BlockProperties{}
}

func (b CutSandstone) New(props BlockProperties) Block {
	return CutSandstone{}
}