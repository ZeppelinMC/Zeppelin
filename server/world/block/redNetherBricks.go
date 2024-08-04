package block

type RedNetherBricks struct {
}

func (b RedNetherBricks) Encode() (string, BlockProperties) {
	return "minecraft:red_nether_bricks", BlockProperties{}
}

func (b RedNetherBricks) New(props BlockProperties) Block {
	return RedNetherBricks{}
}