package block

type PottedOakSapling struct {
}

func (b PottedOakSapling) Encode() (string, BlockProperties) {
	return "minecraft:potted_oak_sapling", BlockProperties{}
}

func (b PottedOakSapling) New(props BlockProperties) Block {
	return PottedOakSapling{}
}