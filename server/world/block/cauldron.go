package block

type Cauldron struct {
}

func (b Cauldron) Encode() (string, BlockProperties) {
	return "minecraft:cauldron", BlockProperties{}
}

func (b Cauldron) New(props BlockProperties) Block {
	return Cauldron{}
}