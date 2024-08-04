package block

type PottedFloweringAzaleaBush struct {
}

func (b PottedFloweringAzaleaBush) Encode() (string, BlockProperties) {
	return "minecraft:potted_flowering_azalea_bush", BlockProperties{}
}

func (b PottedFloweringAzaleaBush) New(props BlockProperties) Block {
	return PottedFloweringAzaleaBush{}
}