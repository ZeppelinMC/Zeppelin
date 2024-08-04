package block

type RedSand struct {
}

func (b RedSand) Encode() (string, BlockProperties) {
	return "minecraft:red_sand", BlockProperties{}
}

func (b RedSand) New(props BlockProperties) Block {
	return RedSand{}
}