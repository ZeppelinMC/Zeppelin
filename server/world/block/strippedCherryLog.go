package block

type StrippedCherryLog struct {
	Axis string
}

func (b StrippedCherryLog) Encode() (string, BlockProperties) {
	return "minecraft:stripped_cherry_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedCherryLog) New(props BlockProperties) Block {
	return StrippedCherryLog{
		Axis: props["axis"],
	}
}