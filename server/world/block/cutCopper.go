package block

type CutCopper struct {
}

func (b CutCopper) Encode() (string, BlockProperties) {
	return "minecraft:cut_copper", BlockProperties{}
}

func (b CutCopper) New(props BlockProperties) Block {
	return CutCopper{}
}