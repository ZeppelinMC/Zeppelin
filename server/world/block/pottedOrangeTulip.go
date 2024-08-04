package block

type PottedOrangeTulip struct {
}

func (b PottedOrangeTulip) Encode() (string, BlockProperties) {
	return "minecraft:potted_orange_tulip", BlockProperties{}
}

func (b PottedOrangeTulip) New(props BlockProperties) Block {
	return PottedOrangeTulip{}
}