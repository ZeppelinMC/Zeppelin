package block

type PottedCornflower struct {
}

func (b PottedCornflower) Encode() (string, BlockProperties) {
	return "minecraft:potted_cornflower", BlockProperties{}
}

func (b PottedCornflower) New(props BlockProperties) Block {
	return PottedCornflower{}
}