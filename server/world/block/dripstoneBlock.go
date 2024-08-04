package block

type DripstoneBlock struct {
}

func (b DripstoneBlock) Encode() (string, BlockProperties) {
	return "minecraft:dripstone_block", BlockProperties{}
}

func (b DripstoneBlock) New(props BlockProperties) Block {
	return DripstoneBlock{}
}