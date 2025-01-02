package block

type RedStainedGlass struct {
}

func (b RedStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:red_stained_glass", BlockProperties{}
}

func (b RedStainedGlass) New(props BlockProperties) Block {
	return RedStainedGlass{}
}