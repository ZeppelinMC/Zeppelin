package block

type OxidizedCopper struct {
}

func (b OxidizedCopper) Encode() (string, BlockProperties) {
	return "minecraft:oxidized_copper", BlockProperties{}
}

func (b OxidizedCopper) New(props BlockProperties) Block {
	return OxidizedCopper{}
}