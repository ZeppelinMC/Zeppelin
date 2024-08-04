package block

type ExposedCutCopper struct {
}

func (b ExposedCutCopper) Encode() (string, BlockProperties) {
	return "minecraft:exposed_cut_copper", BlockProperties{}
}

func (b ExposedCutCopper) New(props BlockProperties) Block {
	return ExposedCutCopper{}
}