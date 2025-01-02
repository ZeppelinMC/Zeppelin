package block

type MangroveLog struct {
	Axis string
}

func (b MangroveLog) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b MangroveLog) New(props BlockProperties) Block {
	return MangroveLog{
		Axis: props["axis"],
	}
}