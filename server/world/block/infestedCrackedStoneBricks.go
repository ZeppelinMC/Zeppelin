package block

type InfestedCrackedStoneBricks struct {
}

func (b InfestedCrackedStoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:infested_cracked_stone_bricks", BlockProperties{}
}

func (b InfestedCrackedStoneBricks) New(props BlockProperties) Block {
	return InfestedCrackedStoneBricks{}
}