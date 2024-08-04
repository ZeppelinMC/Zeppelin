package block

type RedstoneBlock struct {
}

func (b RedstoneBlock) Encode() (string, BlockProperties) {
	return "minecraft:redstone_block", BlockProperties{}
}

func (b RedstoneBlock) New(props BlockProperties) Block {
	return RedstoneBlock{}
}