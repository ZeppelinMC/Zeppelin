package block

type Deepslate struct {
	Axis string
}

func (b Deepslate) Encode() (string, BlockProperties) {
	return "minecraft:deepslate", BlockProperties{
		"axis": b.Axis,
	}
}

func (b Deepslate) New(props BlockProperties) Block {
	return Deepslate{
		Axis: props["axis"],
	}
}
