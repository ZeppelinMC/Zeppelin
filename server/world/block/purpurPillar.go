package block

type PurpurPillar struct {
	Axis string
}

func (b PurpurPillar) Encode() (string, BlockProperties) {
	return "minecraft:purpur_pillar", BlockProperties{
		"axis": b.Axis,
	}
}

func (b PurpurPillar) New(props BlockProperties) Block {
	return PurpurPillar{
		Axis: props["axis"],
	}
}