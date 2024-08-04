package block

type RedMushroom struct {
}

func (b RedMushroom) Encode() (string, BlockProperties) {
	return "minecraft:red_mushroom", BlockProperties{}
}

func (b RedMushroom) New(props BlockProperties) Block {
	return RedMushroom{}
}