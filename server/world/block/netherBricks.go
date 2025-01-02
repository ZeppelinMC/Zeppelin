package block

type NetherBricks struct {
}

func (b NetherBricks) Encode() (string, BlockProperties) {
	return "minecraft:nether_bricks", BlockProperties{}
}

func (b NetherBricks) New(props BlockProperties) Block {
	return NetherBricks{}
}