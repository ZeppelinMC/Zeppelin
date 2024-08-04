package block

type WaxedOxidizedCopper struct {
}

func (b WaxedOxidizedCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_oxidized_copper", BlockProperties{}
}

func (b WaxedOxidizedCopper) New(props BlockProperties) Block {
	return WaxedOxidizedCopper{}
}