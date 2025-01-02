package block

type PottedAcaciaSapling struct {
}

func (b PottedAcaciaSapling) Encode() (string, BlockProperties) {
	return "minecraft:potted_acacia_sapling", BlockProperties{}
}

func (b PottedAcaciaSapling) New(props BlockProperties) Block {
	return PottedAcaciaSapling{}
}