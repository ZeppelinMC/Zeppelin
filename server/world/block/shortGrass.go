package block

type ShortGrass struct {
}

func (b ShortGrass) Encode() (string, BlockProperties) {
	return "minecraft:short_grass", BlockProperties{}
}

func (b ShortGrass) New(props BlockProperties) Block {
	return ShortGrass{}
}