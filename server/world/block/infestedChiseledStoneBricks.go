package block

type InfestedChiseledStoneBricks struct {
}

func (b InfestedChiseledStoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:infested_chiseled_stone_bricks", BlockProperties{}
}

func (b InfestedChiseledStoneBricks) New(props BlockProperties) Block {
	return InfestedChiseledStoneBricks{}
}