package block

type PottedCherrySapling struct {
}

func (b PottedCherrySapling) Encode() (string, BlockProperties) {
	return "minecraft:potted_cherry_sapling", BlockProperties{}
}

func (b PottedCherrySapling) New(props BlockProperties) Block {
	return PottedCherrySapling{}
}