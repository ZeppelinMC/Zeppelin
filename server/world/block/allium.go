package block

type Allium struct {
}

func (b Allium) Encode() (string, BlockProperties) {
	return "minecraft:allium", BlockProperties{}
}

func (b Allium) New(props BlockProperties) Block {
	return Allium{}
}