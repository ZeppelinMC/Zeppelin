package block

type CrimsonStem struct {
	Axis string
}

func (b CrimsonStem) Encode() (string, BlockProperties) {
	return "minecraft:crimson_stem", BlockProperties{
		"axis": b.Axis,
	}
}

func (b CrimsonStem) New(props BlockProperties) Block {
	return CrimsonStem{
		Axis: props["axis"],
	}
}