package block

type FireCoralBlock struct {
}

func (b FireCoralBlock) Encode() (string, BlockProperties) {
	return "minecraft:fire_coral_block", BlockProperties{}
}

func (b FireCoralBlock) New(props BlockProperties) Block {
	return FireCoralBlock{}
}