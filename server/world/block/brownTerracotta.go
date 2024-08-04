package block

type BrownTerracotta struct {
}

func (b BrownTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:brown_terracotta", BlockProperties{}
}

func (b BrownTerracotta) New(props BlockProperties) Block {
	return BrownTerracotta{}
}