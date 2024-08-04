package block

type LimeCarpet struct {
}

func (b LimeCarpet) Encode() (string, BlockProperties) {
	return "minecraft:lime_carpet", BlockProperties{}
}

func (b LimeCarpet) New(props BlockProperties) Block {
	return LimeCarpet{}
}