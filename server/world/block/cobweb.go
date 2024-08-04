package block

type Cobweb struct {
}

func (b Cobweb) Encode() (string, BlockProperties) {
	return "minecraft:cobweb", BlockProperties{}
}

func (b Cobweb) New(props BlockProperties) Block {
	return Cobweb{}
}