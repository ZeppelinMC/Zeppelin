package block

type WeepingVinesPlant struct {
}

func (b WeepingVinesPlant) Encode() (string, BlockProperties) {
	return "minecraft:weeping_vines_plant", BlockProperties{}
}

func (b WeepingVinesPlant) New(props BlockProperties) Block {
	return WeepingVinesPlant{}
}