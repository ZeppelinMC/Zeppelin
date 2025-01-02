package block

type Lilac struct {
	Half string
}

func (b Lilac) Encode() (string, BlockProperties) {
	return "minecraft:lilac", BlockProperties{
		"half": b.Half,
	}
}

func (b Lilac) New(props BlockProperties) Block {
	return Lilac{
		Half: props["half"],
	}
}