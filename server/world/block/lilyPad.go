package block

type LilyPad struct {
}

func (b LilyPad) Encode() (string, BlockProperties) {
	return "minecraft:lily_pad", BlockProperties{}
}

func (b LilyPad) New(props BlockProperties) Block {
	return LilyPad{}
}