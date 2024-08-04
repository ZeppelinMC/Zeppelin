package block

type JungleLog struct {
	Axis string
}

func (b JungleLog) Encode() (string, BlockProperties) {
	return "minecraft:jungle_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b JungleLog) New(props BlockProperties) Block {
	return JungleLog{
		Axis: props["axis"],
	}
}