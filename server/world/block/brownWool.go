package block

type BrownWool struct {
}

func (b BrownWool) Encode() (string, BlockProperties) {
	return "minecraft:brown_wool", BlockProperties{}
}

func (b BrownWool) New(props BlockProperties) Block {
	return BrownWool{}
}