package block

type WhiteTulip struct {
}

func (b WhiteTulip) Encode() (string, BlockProperties) {
	return "minecraft:white_tulip", BlockProperties{}
}

func (b WhiteTulip) New(props BlockProperties) Block {
	return WhiteTulip{}
}