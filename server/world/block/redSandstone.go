package block

type RedSandstone struct {
}

func (b RedSandstone) Encode() (string, BlockProperties) {
	return "minecraft:red_sandstone", BlockProperties{}
}

func (b RedSandstone) New(props BlockProperties) Block {
	return RedSandstone{}
}