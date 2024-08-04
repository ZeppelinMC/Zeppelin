package block

type HoneycombBlock struct {
}

func (b HoneycombBlock) Encode() (string, BlockProperties) {
	return "minecraft:honeycomb_block", BlockProperties{}
}

func (b HoneycombBlock) New(props BlockProperties) Block {
	return HoneycombBlock{}
}