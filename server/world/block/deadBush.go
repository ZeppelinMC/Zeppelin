package block

type DeadBush struct {
}

func (b DeadBush) Encode() (string, BlockProperties) {
	return "minecraft:dead_bush", BlockProperties{}
}

func (b DeadBush) New(props BlockProperties) Block {
	return DeadBush{}
}