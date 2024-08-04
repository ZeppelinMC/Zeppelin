package block

type MagentaConcretePowder struct {
}

func (b MagentaConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:magenta_concrete_powder", BlockProperties{}
}

func (b MagentaConcretePowder) New(props BlockProperties) Block {
	return MagentaConcretePowder{}
}