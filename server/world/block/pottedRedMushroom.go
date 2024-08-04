package block

type PottedRedMushroom struct {
}

func (b PottedRedMushroom) Encode() (string, BlockProperties) {
	return "minecraft:potted_red_mushroom", BlockProperties{}
}

func (b PottedRedMushroom) New(props BlockProperties) Block {
	return PottedRedMushroom{}
}