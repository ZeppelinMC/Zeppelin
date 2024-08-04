package block

type WaxedCutCopper struct {
}

func (b WaxedCutCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_cut_copper", BlockProperties{}
}

func (b WaxedCutCopper) New(props BlockProperties) Block {
	return WaxedCutCopper{}
}