package block

type Netherrack struct {
}

func (b Netherrack) Encode() (string, BlockProperties) {
	return "minecraft:netherrack", BlockProperties{}
}

func (b Netherrack) New(props BlockProperties) Block {
	return Netherrack{}
}