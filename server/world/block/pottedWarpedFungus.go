package block

type PottedWarpedFungus struct {
}

func (b PottedWarpedFungus) Encode() (string, BlockProperties) {
	return "minecraft:potted_warped_fungus", BlockProperties{}
}

func (b PottedWarpedFungus) New(props BlockProperties) Block {
	return PottedWarpedFungus{}
}