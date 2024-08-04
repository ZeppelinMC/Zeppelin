package block

type Sunflower struct {
	Half string
}

func (b Sunflower) Encode() (string, BlockProperties) {
	return "minecraft:sunflower", BlockProperties{
		"half": b.Half,
	}
}

func (b Sunflower) New(props BlockProperties) Block {
	return Sunflower{
		Half: props["half"],
	}
}