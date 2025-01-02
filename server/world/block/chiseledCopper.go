package block

type ChiseledCopper struct {
}

func (b ChiseledCopper) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_copper", BlockProperties{}
}

func (b ChiseledCopper) New(props BlockProperties) Block {
	return ChiseledCopper{}
}