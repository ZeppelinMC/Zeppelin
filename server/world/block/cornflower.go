package block

type Cornflower struct {
}

func (b Cornflower) Encode() (string, BlockProperties) {
	return "minecraft:cornflower", BlockProperties{}
}

func (b Cornflower) New(props BlockProperties) Block {
	return Cornflower{}
}