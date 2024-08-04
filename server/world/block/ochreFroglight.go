package block

type OchreFroglight struct {
	Axis string
}

func (b OchreFroglight) Encode() (string, BlockProperties) {
	return "minecraft:ochre_froglight", BlockProperties{
		"axis": b.Axis,
	}
}

func (b OchreFroglight) New(props BlockProperties) Block {
	return OchreFroglight{
		Axis: props["axis"],
	}
}