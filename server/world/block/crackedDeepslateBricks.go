package block

type CrackedDeepslateBricks struct {
}

func (b CrackedDeepslateBricks) Encode() (string, BlockProperties) {
	return "minecraft:cracked_deepslate_bricks", BlockProperties{}
}

func (b CrackedDeepslateBricks) New(props BlockProperties) Block {
	return CrackedDeepslateBricks{}
}