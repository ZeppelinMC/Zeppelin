package block

type PolishedDiorite struct {
}

func (b PolishedDiorite) Encode() (string, BlockProperties) {
	return "minecraft:polished_diorite", BlockProperties{}
}

func (b PolishedDiorite) New(props BlockProperties) Block {
	return PolishedDiorite{}
}