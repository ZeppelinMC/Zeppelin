package block

type InfestedStoneBricks struct {
}

func (b InfestedStoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:infested_stone_bricks", BlockProperties{}
}

func (b InfestedStoneBricks) New(props BlockProperties) Block {
	return InfestedStoneBricks{}
}