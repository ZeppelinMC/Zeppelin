package block

type TallGrass struct {
	Half string
}

func (b TallGrass) Encode() (string, BlockProperties) {
	return "minecraft:tall_grass", BlockProperties{
		"half": b.Half,
	}
}

func (b TallGrass) New(props BlockProperties) Block {
	return TallGrass{
		Half: props["half"],
	}
}