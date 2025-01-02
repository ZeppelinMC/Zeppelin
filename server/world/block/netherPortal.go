package block

type NetherPortal struct {
	Axis string
}

func (b NetherPortal) Encode() (string, BlockProperties) {
	return "minecraft:nether_portal", BlockProperties{
		"axis": b.Axis,
	}
}

func (b NetherPortal) New(props BlockProperties) Block {
	return NetherPortal{
		Axis: props["axis"],
	}
}