package block

type Fern struct {
}

func (b Fern) Encode() (string, BlockProperties) {
	return "minecraft:fern", BlockProperties{}
}

func (b Fern) New(props BlockProperties) Block {
	return Fern{}
}