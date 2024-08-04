package block

type BambooSapling struct {
}

func (b BambooSapling) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_sapling", BlockProperties{}
}

func (b BambooSapling) New(props BlockProperties) Block {
	return BambooSapling{}
}