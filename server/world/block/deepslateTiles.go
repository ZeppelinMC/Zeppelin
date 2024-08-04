package block

type DeepslateTiles struct {
}

func (b DeepslateTiles) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_tiles", BlockProperties{}
}

func (b DeepslateTiles) New(props BlockProperties) Block {
	return DeepslateTiles{}
}