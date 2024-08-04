package block

type Sand struct {
}

func (b Sand) Encode() (string, BlockProperties) {
	return "minecraft:sand", BlockProperties{}
}

func (b Sand) New(props BlockProperties) Block {
	return Sand{}
}