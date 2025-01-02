package block

type WaxedOxidizedChiseledCopper struct {
}

func (b WaxedOxidizedChiseledCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_oxidized_chiseled_copper", BlockProperties{}
}

func (b WaxedOxidizedChiseledCopper) New(props BlockProperties) Block {
	return WaxedOxidizedChiseledCopper{}
}