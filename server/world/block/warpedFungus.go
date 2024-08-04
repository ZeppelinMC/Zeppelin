package block

type WarpedFungus struct {
}

func (b WarpedFungus) Encode() (string, BlockProperties) {
	return "minecraft:warped_fungus", BlockProperties{}
}

func (b WarpedFungus) New(props BlockProperties) Block {
	return WarpedFungus{}
}