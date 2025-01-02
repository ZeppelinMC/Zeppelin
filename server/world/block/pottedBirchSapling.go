package block

type PottedBirchSapling struct {
}

func (b PottedBirchSapling) Encode() (string, BlockProperties) {
	return "minecraft:potted_birch_sapling", BlockProperties{}
}

func (b PottedBirchSapling) New(props BlockProperties) Block {
	return PottedBirchSapling{}
}