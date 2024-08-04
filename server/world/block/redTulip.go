package block

type RedTulip struct {
}

func (b RedTulip) Encode() (string, BlockProperties) {
	return "minecraft:red_tulip", BlockProperties{}
}

func (b RedTulip) New(props BlockProperties) Block {
	return RedTulip{}
}