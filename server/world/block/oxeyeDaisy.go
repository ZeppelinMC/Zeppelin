package block

type OxeyeDaisy struct {
}

func (b OxeyeDaisy) Encode() (string, BlockProperties) {
	return "minecraft:oxeye_daisy", BlockProperties{}
}

func (b OxeyeDaisy) New(props BlockProperties) Block {
	return OxeyeDaisy{}
}