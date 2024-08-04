package block

type CrackedStoneBricks struct {
}

func (b CrackedStoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:cracked_stone_bricks", BlockProperties{}
}

func (b CrackedStoneBricks) New(props BlockProperties) Block {
	return CrackedStoneBricks{}
}