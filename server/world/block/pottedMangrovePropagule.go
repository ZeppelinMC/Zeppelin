package block

type PottedMangrovePropagule struct {
}

func (b PottedMangrovePropagule) Encode() (string, BlockProperties) {
	return "minecraft:potted_mangrove_propagule", BlockProperties{}
}

func (b PottedMangrovePropagule) New(props BlockProperties) Block {
	return PottedMangrovePropagule{}
}