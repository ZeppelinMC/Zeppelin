package block

type CoalBlock struct {
}

func (b CoalBlock) Encode() (string, BlockProperties) {
	return "minecraft:coal_block", BlockProperties{}
}

func (b CoalBlock) New(props BlockProperties) Block {
	return CoalBlock{}
}