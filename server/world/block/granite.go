package block

type Granite struct {
}

func (b Granite) Encode() (string, BlockProperties) {
	return "minecraft:granite", BlockProperties{}
}

func (b Granite) New(props BlockProperties) Block {
	return Granite{}
}