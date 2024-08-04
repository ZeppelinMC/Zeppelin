package block

type DeadBubbleCoralBlock struct {
}

func (b DeadBubbleCoralBlock) Encode() (string, BlockProperties) {
	return "minecraft:dead_bubble_coral_block", BlockProperties{}
}

func (b DeadBubbleCoralBlock) New(props BlockProperties) Block {
	return DeadBubbleCoralBlock{}
}