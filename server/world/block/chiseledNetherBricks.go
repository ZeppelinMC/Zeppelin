package block

type ChiseledNetherBricks struct {
}

func (b ChiseledNetherBricks) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_nether_bricks", BlockProperties{}
}

func (b ChiseledNetherBricks) New(props BlockProperties) Block {
	return ChiseledNetherBricks{}
}