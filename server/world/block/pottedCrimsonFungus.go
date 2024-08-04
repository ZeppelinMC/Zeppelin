package block

type PottedCrimsonFungus struct {
}

func (b PottedCrimsonFungus) Encode() (string, BlockProperties) {
	return "minecraft:potted_crimson_fungus", BlockProperties{}
}

func (b PottedCrimsonFungus) New(props BlockProperties) Block {
	return PottedCrimsonFungus{}
}