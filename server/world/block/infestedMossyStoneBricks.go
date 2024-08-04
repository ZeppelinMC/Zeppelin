package block

type InfestedMossyStoneBricks struct {
}

func (b InfestedMossyStoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:infested_mossy_stone_bricks", BlockProperties{}
}

func (b InfestedMossyStoneBricks) New(props BlockProperties) Block {
	return InfestedMossyStoneBricks{}
}