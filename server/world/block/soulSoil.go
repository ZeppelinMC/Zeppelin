package block

type SoulSoil struct {
}

func (b SoulSoil) Encode() (string, BlockProperties) {
	return "minecraft:soul_soil", BlockProperties{}
}

func (b SoulSoil) New(props BlockProperties) Block {
	return SoulSoil{}
}