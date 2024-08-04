package block

type Diorite struct {
}

func (b Diorite) Encode() (string, BlockProperties) {
	return "minecraft:diorite", BlockProperties{}
}

func (b Diorite) New(props BlockProperties) Block {
	return Diorite{}
}