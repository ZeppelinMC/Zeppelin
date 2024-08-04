package block

type Torch struct {
}

func (b Torch) Encode() (string, BlockProperties) {
	return "minecraft:torch", BlockProperties{}
}

func (b Torch) New(props BlockProperties) Block {
	return Torch{}
}