package block

type WeatheredChiseledCopper struct {
}

func (b WeatheredChiseledCopper) Encode() (string, BlockProperties) {
	return "minecraft:weathered_chiseled_copper", BlockProperties{}
}

func (b WeatheredChiseledCopper) New(props BlockProperties) Block {
	return WeatheredChiseledCopper{}
}