package block

type OxidizedCutCopper struct {
}

func (b OxidizedCutCopper) Encode() (string, BlockProperties) {
	return "minecraft:oxidized_cut_copper", BlockProperties{}
}

func (b OxidizedCutCopper) New(props BlockProperties) Block {
	return OxidizedCutCopper{}
}