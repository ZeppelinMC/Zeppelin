package block

type Sculk struct {
}

func (b Sculk) Encode() (string, BlockProperties) {
	return "minecraft:sculk", BlockProperties{}
}

func (b Sculk) New(props BlockProperties) Block {
	return Sculk{}
}