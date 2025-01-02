package block

type DeadBrainCoralBlock struct {
}

func (b DeadBrainCoralBlock) Encode() (string, BlockProperties) {
	return "minecraft:dead_brain_coral_block", BlockProperties{}
}

func (b DeadBrainCoralBlock) New(props BlockProperties) Block {
	return DeadBrainCoralBlock{}
}