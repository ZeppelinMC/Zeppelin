package block

type MossyStoneBricks struct {
}

func (b MossyStoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:mossy_stone_bricks", BlockProperties{}
}

func (b MossyStoneBricks) New(props BlockProperties) Block {
	return MossyStoneBricks{}
}