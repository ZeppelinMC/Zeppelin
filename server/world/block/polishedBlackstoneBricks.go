package block

type PolishedBlackstoneBricks struct {
}

func (b PolishedBlackstoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:polished_blackstone_bricks", BlockProperties{}
}

func (b PolishedBlackstoneBricks) New(props BlockProperties) Block {
	return PolishedBlackstoneBricks{}
}