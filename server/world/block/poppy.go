package block

type Poppy struct {
}

func (b Poppy) Encode() (string, BlockProperties) {
	return "minecraft:poppy", BlockProperties{}
}

func (b Poppy) New(props BlockProperties) Block {
	return Poppy{}
}