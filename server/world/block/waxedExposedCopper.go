package block

type WaxedExposedCopper struct {
}

func (b WaxedExposedCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_exposed_copper", BlockProperties{}
}

func (b WaxedExposedCopper) New(props BlockProperties) Block {
	return WaxedExposedCopper{}
}