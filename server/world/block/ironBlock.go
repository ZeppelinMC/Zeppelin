package block

type IronBlock struct {
}

func (b IronBlock) Encode() (string, BlockProperties) {
	return "minecraft:iron_block", BlockProperties{}
}

func (b IronBlock) New(props BlockProperties) Block {
	return IronBlock{}
}