package block

type PottedWitherRose struct {
}

func (b PottedWitherRose) Encode() (string, BlockProperties) {
	return "minecraft:potted_wither_rose", BlockProperties{}
}

func (b PottedWitherRose) New(props BlockProperties) Block {
	return PottedWitherRose{}
}