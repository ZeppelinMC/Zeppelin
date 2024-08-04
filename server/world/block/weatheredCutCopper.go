package block

type WeatheredCutCopper struct {
}

func (b WeatheredCutCopper) Encode() (string, BlockProperties) {
	return "minecraft:weathered_cut_copper", BlockProperties{}
}

func (b WeatheredCutCopper) New(props BlockProperties) Block {
	return WeatheredCutCopper{}
}