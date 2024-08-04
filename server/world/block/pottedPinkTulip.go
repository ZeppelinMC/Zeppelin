package block

type PottedPinkTulip struct {
}

func (b PottedPinkTulip) Encode() (string, BlockProperties) {
	return "minecraft:potted_pink_tulip", BlockProperties{}
}

func (b PottedPinkTulip) New(props BlockProperties) Block {
	return PottedPinkTulip{}
}