package block

type WaxedOxidizedCutCopper struct {
}

func (b WaxedOxidizedCutCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_oxidized_cut_copper", BlockProperties{}
}

func (b WaxedOxidizedCutCopper) New(props BlockProperties) Block {
	return WaxedOxidizedCutCopper{}
}