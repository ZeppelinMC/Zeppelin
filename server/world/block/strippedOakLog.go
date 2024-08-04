package block

type StrippedOakLog struct {
	Axis string
}

func (b StrippedOakLog) Encode() (string, BlockProperties) {
	return "minecraft:stripped_oak_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedOakLog) New(props BlockProperties) Block {
	return StrippedOakLog{
		Axis: props["axis"],
	}
}