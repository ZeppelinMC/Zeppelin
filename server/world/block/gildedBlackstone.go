package block

type GildedBlackstone struct {
}

func (b GildedBlackstone) Encode() (string, BlockProperties) {
	return "minecraft:gilded_blackstone", BlockProperties{}
}

func (b GildedBlackstone) New(props BlockProperties) Block {
	return GildedBlackstone{}
}