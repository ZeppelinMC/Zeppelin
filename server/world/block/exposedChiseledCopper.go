package block

type ExposedChiseledCopper struct {
}

func (b ExposedChiseledCopper) Encode() (string, BlockProperties) {
	return "minecraft:exposed_chiseled_copper", BlockProperties{}
}

func (b ExposedChiseledCopper) New(props BlockProperties) Block {
	return ExposedChiseledCopper{}
}