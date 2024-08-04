package block

type Grindstone struct {
	Face string
	Facing string
}

func (b Grindstone) Encode() (string, BlockProperties) {
	return "minecraft:grindstone", BlockProperties{
		"face": b.Face,
		"facing": b.Facing,
	}
}

func (b Grindstone) New(props BlockProperties) Block {
	return Grindstone{
		Face: props["face"],
		Facing: props["facing"],
	}
}