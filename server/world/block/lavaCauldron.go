package block

type LavaCauldron struct {
}

func (b LavaCauldron) Encode() (string, BlockProperties) {
	return "minecraft:lava_cauldron", BlockProperties{}
}

func (b LavaCauldron) New(props BlockProperties) Block {
	return LavaCauldron{}
}