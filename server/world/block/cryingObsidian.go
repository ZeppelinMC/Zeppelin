package block

type CryingObsidian struct {
}

func (b CryingObsidian) Encode() (string, BlockProperties) {
	return "minecraft:crying_obsidian", BlockProperties{}
}

func (b CryingObsidian) New(props BlockProperties) Block {
	return CryingObsidian{}
}