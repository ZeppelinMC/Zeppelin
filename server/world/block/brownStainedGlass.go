package block

type BrownStainedGlass struct {
}

func (b BrownStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:brown_stained_glass", BlockProperties{}
}

func (b BrownStainedGlass) New(props BlockProperties) Block {
	return BrownStainedGlass{}
}