package block

type PottedOxeyeDaisy struct {
}

func (b PottedOxeyeDaisy) Encode() (string, BlockProperties) {
	return "minecraft:potted_oxeye_daisy", BlockProperties{}
}

func (b PottedOxeyeDaisy) New(props BlockProperties) Block {
	return PottedOxeyeDaisy{}
}