package block

type OrangeCarpet struct {
}

func (b OrangeCarpet) Encode() (string, BlockProperties) {
	return "minecraft:orange_carpet", BlockProperties{}
}

func (b OrangeCarpet) New(props BlockProperties) Block {
	return OrangeCarpet{}
}