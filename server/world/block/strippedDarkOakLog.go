package block

type StrippedDarkOakLog struct {
	Axis string
}

func (b StrippedDarkOakLog) Encode() (string, BlockProperties) {
	return "minecraft:stripped_dark_oak_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedDarkOakLog) New(props BlockProperties) Block {
	return StrippedDarkOakLog{
		Axis: props["axis"],
	}
}