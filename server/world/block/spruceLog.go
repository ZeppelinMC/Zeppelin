package block

type SpruceLog struct {
	Axis string
}

func (b SpruceLog) Encode() (string, BlockProperties) {
	return "minecraft:spruce_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b SpruceLog) New(props BlockProperties) Block {
	return SpruceLog{
		Axis: props["axis"],
	}
}