package block

type OakLog struct {
	Axis string
}

func (b OakLog) Encode() (string, BlockProperties) {
	return "minecraft:oak_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b OakLog) New(props BlockProperties) Block {
	return OakLog{
		Axis: props["axis"],
	}
}