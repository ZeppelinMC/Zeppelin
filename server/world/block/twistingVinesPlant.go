package block

type TwistingVinesPlant struct {
}

func (b TwistingVinesPlant) Encode() (string, BlockProperties) {
	return "minecraft:twisting_vines_plant", BlockProperties{}
}

func (b TwistingVinesPlant) New(props BlockProperties) Block {
	return TwistingVinesPlant{}
}