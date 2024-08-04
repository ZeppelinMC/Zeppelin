package block

type Torchflower struct {
}

func (b Torchflower) Encode() (string, BlockProperties) {
	return "minecraft:torchflower", BlockProperties{}
}

func (b Torchflower) New(props BlockProperties) Block {
	return Torchflower{}
}