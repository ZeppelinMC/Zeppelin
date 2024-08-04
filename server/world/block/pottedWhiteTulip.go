package block

type PottedWhiteTulip struct {
}

func (b PottedWhiteTulip) Encode() (string, BlockProperties) {
	return "minecraft:potted_white_tulip", BlockProperties{}
}

func (b PottedWhiteTulip) New(props BlockProperties) Block {
	return PottedWhiteTulip{}
}