package block

type MagentaConcrete struct {
}

func (b MagentaConcrete) Encode() (string, BlockProperties) {
	return "minecraft:magenta_concrete", BlockProperties{}
}

func (b MagentaConcrete) New(props BlockProperties) Block {
	return MagentaConcrete{}
}