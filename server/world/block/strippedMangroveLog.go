package block

type StrippedMangroveLog struct {
	Axis string
}

func (b StrippedMangroveLog) Encode() (string, BlockProperties) {
	return "minecraft:stripped_mangrove_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedMangroveLog) New(props BlockProperties) Block {
	return StrippedMangroveLog{
		Axis: props["axis"],
	}
}