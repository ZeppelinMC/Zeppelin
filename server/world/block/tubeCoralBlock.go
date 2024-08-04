package block

type TubeCoralBlock struct {
}

func (b TubeCoralBlock) Encode() (string, BlockProperties) {
	return "minecraft:tube_coral_block", BlockProperties{}
}

func (b TubeCoralBlock) New(props BlockProperties) Block {
	return TubeCoralBlock{}
}