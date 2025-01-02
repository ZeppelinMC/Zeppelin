package block

type PottedDeadBush struct {
}

func (b PottedDeadBush) Encode() (string, BlockProperties) {
	return "minecraft:potted_dead_bush", BlockProperties{}
}

func (b PottedDeadBush) New(props BlockProperties) Block {
	return PottedDeadBush{}
}