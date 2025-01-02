package block

type DeadFireCoralBlock struct {
}

func (b DeadFireCoralBlock) Encode() (string, BlockProperties) {
	return "minecraft:dead_fire_coral_block", BlockProperties{}
}

func (b DeadFireCoralBlock) New(props BlockProperties) Block {
	return DeadFireCoralBlock{}
}