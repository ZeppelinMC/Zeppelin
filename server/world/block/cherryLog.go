package block

type CherryLog struct {
	Axis string
}

func (b CherryLog) Encode() (string, BlockProperties) {
	return "minecraft:cherry_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b CherryLog) New(props BlockProperties) Block {
	return CherryLog{
		Axis: props["axis"],
	}
}