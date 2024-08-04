package block

type BirchLog struct {
	Axis string
}

func (b BirchLog) Encode() (string, BlockProperties) {
	return "minecraft:birch_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b BirchLog) New(props BlockProperties) Block {
	return BirchLog{
		Axis: props["axis"],
	}
}