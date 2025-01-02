package block

type OxidizedChiseledCopper struct {
}

func (b OxidizedChiseledCopper) Encode() (string, BlockProperties) {
	return "minecraft:oxidized_chiseled_copper", BlockProperties{}
}

func (b OxidizedChiseledCopper) New(props BlockProperties) Block {
	return OxidizedChiseledCopper{}
}