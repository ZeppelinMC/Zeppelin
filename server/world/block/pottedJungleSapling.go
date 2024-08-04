package block

type PottedJungleSapling struct {
}

func (b PottedJungleSapling) Encode() (string, BlockProperties) {
	return "minecraft:potted_jungle_sapling", BlockProperties{}
}

func (b PottedJungleSapling) New(props BlockProperties) Block {
	return PottedJungleSapling{}
}