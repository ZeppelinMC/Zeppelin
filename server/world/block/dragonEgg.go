package block

type DragonEgg struct {
}

func (b DragonEgg) Encode() (string, BlockProperties) {
	return "minecraft:dragon_egg", BlockProperties{}
}

func (b DragonEgg) New(props BlockProperties) Block {
	return DragonEgg{}
}