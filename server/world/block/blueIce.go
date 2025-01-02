package block

type BlueIce struct {
}

func (b BlueIce) Encode() (string, BlockProperties) {
	return "minecraft:blue_ice", BlockProperties{}
}

func (b BlueIce) New(props BlockProperties) Block {
	return BlueIce{}
}