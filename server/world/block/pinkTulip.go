package block

type PinkTulip struct {
}

func (b PinkTulip) Encode() (string, BlockProperties) {
	return "minecraft:pink_tulip", BlockProperties{}
}

func (b PinkTulip) New(props BlockProperties) Block {
	return PinkTulip{}
}