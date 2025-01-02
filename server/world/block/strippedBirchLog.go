package block

type StrippedBirchLog struct {
	Axis string
}

func (b StrippedBirchLog) Encode() (string, BlockProperties) {
	return "minecraft:stripped_birch_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedBirchLog) New(props BlockProperties) Block {
	return StrippedBirchLog{
		Axis: props["axis"],
	}
}