package block

type CrackedPolishedBlackstoneBricks struct {
}

func (b CrackedPolishedBlackstoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:cracked_polished_blackstone_bricks", BlockProperties{}
}

func (b CrackedPolishedBlackstoneBricks) New(props BlockProperties) Block {
	return CrackedPolishedBlackstoneBricks{}
}