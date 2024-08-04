package block

type ChiseledStoneBricks struct {
}

func (b ChiseledStoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_stone_bricks", BlockProperties{}
}

func (b ChiseledStoneBricks) New(props BlockProperties) Block {
	return ChiseledStoneBricks{}
}