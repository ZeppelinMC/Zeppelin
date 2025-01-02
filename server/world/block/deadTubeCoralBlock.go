package block

type DeadTubeCoralBlock struct {
}

func (b DeadTubeCoralBlock) Encode() (string, BlockProperties) {
	return "minecraft:dead_tube_coral_block", BlockProperties{}
}

func (b DeadTubeCoralBlock) New(props BlockProperties) Block {
	return DeadTubeCoralBlock{}
}