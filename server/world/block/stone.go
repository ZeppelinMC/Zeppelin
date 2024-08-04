package block

type Stone struct {
}

func (b Stone) Encode() (string, BlockProperties) {
	return "minecraft:stone", BlockProperties{}
}

func (b Stone) New(props BlockProperties) Block {
	return Stone{}
}