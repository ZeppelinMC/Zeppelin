package block

type PottedRedTulip struct {
}

func (b PottedRedTulip) Encode() (string, BlockProperties) {
	return "minecraft:potted_red_tulip", BlockProperties{}
}

func (b PottedRedTulip) New(props BlockProperties) Block {
	return PottedRedTulip{}
}