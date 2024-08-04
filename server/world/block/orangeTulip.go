package block

type OrangeTulip struct {
}

func (b OrangeTulip) Encode() (string, BlockProperties) {
	return "minecraft:orange_tulip", BlockProperties{}
}

func (b OrangeTulip) New(props BlockProperties) Block {
	return OrangeTulip{}
}