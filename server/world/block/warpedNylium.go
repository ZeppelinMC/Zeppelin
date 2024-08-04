package block

type WarpedNylium struct {
}

func (b WarpedNylium) Encode() (string, BlockProperties) {
	return "minecraft:warped_nylium", BlockProperties{}
}

func (b WarpedNylium) New(props BlockProperties) Block {
	return WarpedNylium{}
}