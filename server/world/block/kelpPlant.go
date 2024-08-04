package block

type KelpPlant struct {
}

func (b KelpPlant) Encode() (string, BlockProperties) {
	return "minecraft:kelp_plant", BlockProperties{}
}

func (b KelpPlant) New(props BlockProperties) Block {
	return KelpPlant{}
}