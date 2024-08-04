package block

type MagentaCarpet struct {
}

func (b MagentaCarpet) Encode() (string, BlockProperties) {
	return "minecraft:magenta_carpet", BlockProperties{}
}

func (b MagentaCarpet) New(props BlockProperties) Block {
	return MagentaCarpet{}
}