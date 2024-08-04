package block

type PottedAzaleaBush struct {
}

func (b PottedAzaleaBush) Encode() (string, BlockProperties) {
	return "minecraft:potted_azalea_bush", BlockProperties{}
}

func (b PottedAzaleaBush) New(props BlockProperties) Block {
	return PottedAzaleaBush{}
}