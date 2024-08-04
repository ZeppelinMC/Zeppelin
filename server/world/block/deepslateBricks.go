package block

type DeepslateBricks struct {
}

func (b DeepslateBricks) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_bricks", BlockProperties{}
}

func (b DeepslateBricks) New(props BlockProperties) Block {
	return DeepslateBricks{}
}