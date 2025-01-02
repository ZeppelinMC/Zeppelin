package block

type PackedMud struct {
}

func (b PackedMud) Encode() (string, BlockProperties) {
	return "minecraft:packed_mud", BlockProperties{}
}

func (b PackedMud) New(props BlockProperties) Block {
	return PackedMud{}
}