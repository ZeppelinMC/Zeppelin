package block

type WaxedExposedCutCopper struct {
}

func (b WaxedExposedCutCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_exposed_cut_copper", BlockProperties{}
}

func (b WaxedExposedCutCopper) New(props BlockProperties) Block {
	return WaxedExposedCutCopper{}
}