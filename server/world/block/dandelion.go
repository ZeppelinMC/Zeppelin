package block

type Dandelion struct {
}

func (b Dandelion) Encode() (string, BlockProperties) {
	return "minecraft:dandelion", BlockProperties{}
}

func (b Dandelion) New(props BlockProperties) Block {
	return Dandelion{}
}