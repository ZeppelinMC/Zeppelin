package block

type LargeFern struct {
	Half string
}

func (b LargeFern) Encode() (string, BlockProperties) {
	return "minecraft:large_fern", BlockProperties{
		"half": b.Half,
	}
}

func (b LargeFern) New(props BlockProperties) Block {
	return LargeFern{
		Half: props["half"],
	}
}