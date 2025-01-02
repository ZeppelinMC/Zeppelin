package block

type PackedIce struct {
}

func (b PackedIce) Encode() (string, BlockProperties) {
	return "minecraft:packed_ice", BlockProperties{}
}

func (b PackedIce) New(props BlockProperties) Block {
	return PackedIce{}
}