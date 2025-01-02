package block

type QuartzPillar struct {
	Axis string
}

func (b QuartzPillar) Encode() (string, BlockProperties) {
	return "minecraft:quartz_pillar", BlockProperties{
		"axis": b.Axis,
	}
}

func (b QuartzPillar) New(props BlockProperties) Block {
	return QuartzPillar{
		Axis: props["axis"],
	}
}