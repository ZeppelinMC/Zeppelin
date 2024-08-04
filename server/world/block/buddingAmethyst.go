package block

type BuddingAmethyst struct {
}

func (b BuddingAmethyst) Encode() (string, BlockProperties) {
	return "minecraft:budding_amethyst", BlockProperties{}
}

func (b BuddingAmethyst) New(props BlockProperties) Block {
	return BuddingAmethyst{}
}