package block

type StrippedSpruceLog struct {
	Axis string
}

func (b StrippedSpruceLog) Encode() (string, BlockProperties) {
	return "minecraft:stripped_spruce_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedSpruceLog) New(props BlockProperties) Block {
	return StrippedSpruceLog{
		Axis: props["axis"],
	}
}