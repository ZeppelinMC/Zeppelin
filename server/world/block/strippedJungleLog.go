package block

type StrippedJungleLog struct {
	Axis string
}

func (b StrippedJungleLog) Encode() (string, BlockProperties) {
	return "minecraft:stripped_jungle_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedJungleLog) New(props BlockProperties) Block {
	return StrippedJungleLog{
		Axis: props["axis"],
	}
}