package block

type StrippedAcaciaLog struct {
	Axis string
}

func (b StrippedAcaciaLog) Encode() (string, BlockProperties) {
	return "minecraft:stripped_acacia_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedAcaciaLog) New(props BlockProperties) Block {
	return StrippedAcaciaLog{
		Axis: props["axis"],
	}
}