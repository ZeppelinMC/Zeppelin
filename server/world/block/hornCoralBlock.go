package block

type HornCoralBlock struct {
}

func (b HornCoralBlock) Encode() (string, BlockProperties) {
	return "minecraft:horn_coral_block", BlockProperties{}
}

func (b HornCoralBlock) New(props BlockProperties) Block {
	return HornCoralBlock{}
}