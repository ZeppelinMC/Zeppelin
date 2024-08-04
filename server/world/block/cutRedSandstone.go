package block

type CutRedSandstone struct {
}

func (b CutRedSandstone) Encode() (string, BlockProperties) {
	return "minecraft:cut_red_sandstone", BlockProperties{}
}

func (b CutRedSandstone) New(props BlockProperties) Block {
	return CutRedSandstone{}
}