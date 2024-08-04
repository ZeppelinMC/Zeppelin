package block

type PottedDarkOakSapling struct {
}

func (b PottedDarkOakSapling) Encode() (string, BlockProperties) {
	return "minecraft:potted_dark_oak_sapling", BlockProperties{}
}

func (b PottedDarkOakSapling) New(props BlockProperties) Block {
	return PottedDarkOakSapling{}
}