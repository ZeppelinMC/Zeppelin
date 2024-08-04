package block

type ExposedCopper struct {
}

func (b ExposedCopper) Encode() (string, BlockProperties) {
	return "minecraft:exposed_copper", BlockProperties{}
}

func (b ExposedCopper) New(props BlockProperties) Block {
	return ExposedCopper{}
}