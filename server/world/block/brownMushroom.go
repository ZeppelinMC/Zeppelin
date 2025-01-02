package block

type BrownMushroom struct {
}

func (b BrownMushroom) Encode() (string, BlockProperties) {
	return "minecraft:brown_mushroom", BlockProperties{}
}

func (b BrownMushroom) New(props BlockProperties) Block {
	return BrownMushroom{}
}