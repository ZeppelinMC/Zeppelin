package block

type PottedBrownMushroom struct {
}

func (b PottedBrownMushroom) Encode() (string, BlockProperties) {
	return "minecraft:potted_brown_mushroom", BlockProperties{}
}

func (b PottedBrownMushroom) New(props BlockProperties) Block {
	return PottedBrownMushroom{}
}