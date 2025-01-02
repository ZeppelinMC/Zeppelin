package block

type DarkOakLog struct {
	Axis string
}

func (b DarkOakLog) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b DarkOakLog) New(props BlockProperties) Block {
	return DarkOakLog{
		Axis: props["axis"],
	}
}