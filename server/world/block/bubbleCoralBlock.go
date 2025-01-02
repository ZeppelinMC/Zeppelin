package block

type BubbleCoralBlock struct {
}

func (b BubbleCoralBlock) Encode() (string, BlockProperties) {
	return "minecraft:bubble_coral_block", BlockProperties{}
}

func (b BubbleCoralBlock) New(props BlockProperties) Block {
	return BubbleCoralBlock{}
}