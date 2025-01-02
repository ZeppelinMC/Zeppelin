package block

type Lodestone struct {
}

func (b Lodestone) Encode() (string, BlockProperties) {
	return "minecraft:lodestone", BlockProperties{}
}

func (b Lodestone) New(props BlockProperties) Block {
	return Lodestone{}
}