package block

type DeadHornCoralBlock struct {
}

func (b DeadHornCoralBlock) Encode() (string, BlockProperties) {
	return "minecraft:dead_horn_coral_block", BlockProperties{}
}

func (b DeadHornCoralBlock) New(props BlockProperties) Block {
	return DeadHornCoralBlock{}
}