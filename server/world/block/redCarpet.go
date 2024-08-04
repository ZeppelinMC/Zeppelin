package block

type RedCarpet struct {
}

func (b RedCarpet) Encode() (string, BlockProperties) {
	return "minecraft:red_carpet", BlockProperties{}
}

func (b RedCarpet) New(props BlockProperties) Block {
	return RedCarpet{}
}