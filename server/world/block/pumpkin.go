package block

type Pumpkin struct {
}

func (b Pumpkin) Encode() (string, BlockProperties) {
	return "minecraft:pumpkin", BlockProperties{}
}

func (b Pumpkin) New(props BlockProperties) Block {
	return Pumpkin{}
}