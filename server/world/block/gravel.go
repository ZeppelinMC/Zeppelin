package block

type Gravel struct {
}

func (b Gravel) Encode() (string, BlockProperties) {
	return "minecraft:gravel", BlockProperties{}
}

func (b Gravel) New(props BlockProperties) Block {
	return Gravel{}
}