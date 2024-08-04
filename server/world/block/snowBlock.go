package block

type SnowBlock struct {
}

func (b SnowBlock) Encode() (string, BlockProperties) {
	return "minecraft:snow_block", BlockProperties{}
}

func (b SnowBlock) New(props BlockProperties) Block {
	return SnowBlock{}
}