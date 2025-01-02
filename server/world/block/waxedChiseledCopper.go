package block

type WaxedChiseledCopper struct {
}

func (b WaxedChiseledCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_chiseled_copper", BlockProperties{}
}

func (b WaxedChiseledCopper) New(props BlockProperties) Block {
	return WaxedChiseledCopper{}
}