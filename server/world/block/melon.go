package block

type Melon struct {
}

func (b Melon) Encode() (string, BlockProperties) {
	return "minecraft:melon", BlockProperties{}
}

func (b Melon) New(props BlockProperties) Block {
	return Melon{}
}