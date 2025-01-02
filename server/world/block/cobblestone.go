package block

type Cobblestone struct {
}

func (b Cobblestone) Encode() (string, BlockProperties) {
	return "minecraft:cobblestone", BlockProperties{}
}

func (b Cobblestone) New(props BlockProperties) Block {
	return Cobblestone{}
}