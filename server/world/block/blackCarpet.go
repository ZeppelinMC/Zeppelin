package block

type BlackCarpet struct {
}

func (b BlackCarpet) Encode() (string, BlockProperties) {
	return "minecraft:black_carpet", BlockProperties{}
}

func (b BlackCarpet) New(props BlockProperties) Block {
	return BlackCarpet{}
}