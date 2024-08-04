package block

type Peony struct {
	Half string
}

func (b Peony) Encode() (string, BlockProperties) {
	return "minecraft:peony", BlockProperties{
		"half": b.Half,
	}
}

func (b Peony) New(props BlockProperties) Block {
	return Peony{
		Half: props["half"],
	}
}