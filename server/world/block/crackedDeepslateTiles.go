package block

type CrackedDeepslateTiles struct {
}

func (b CrackedDeepslateTiles) Encode() (string, BlockProperties) {
	return "minecraft:cracked_deepslate_tiles", BlockProperties{}
}

func (b CrackedDeepslateTiles) New(props BlockProperties) Block {
	return CrackedDeepslateTiles{}
}