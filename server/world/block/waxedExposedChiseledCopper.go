package block

type WaxedExposedChiseledCopper struct {
}

func (b WaxedExposedChiseledCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_exposed_chiseled_copper", BlockProperties{}
}

func (b WaxedExposedChiseledCopper) New(props BlockProperties) Block {
	return WaxedExposedChiseledCopper{}
}