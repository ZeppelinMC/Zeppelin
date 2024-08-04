package block

type CrackedNetherBricks struct {
}

func (b CrackedNetherBricks) Encode() (string, BlockProperties) {
	return "minecraft:cracked_nether_bricks", BlockProperties{}
}

func (b CrackedNetherBricks) New(props BlockProperties) Block {
	return CrackedNetherBricks{}
}