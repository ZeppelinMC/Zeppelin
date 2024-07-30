package block

type OakLog struct {
	Axis Axis
}

func (g OakLog) Encode() (string, BlockProperties) {
	return "minecraft:oak_log", BlockProperties{
		"axis": g.Axis,
	}
}

func (g OakLog) New(props BlockProperties) Block {
	return OakLog{Axis: props["axis"]}
}

var _ Block = (*OakLog)(nil)
