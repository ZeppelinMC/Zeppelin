package block

type MuddyMangroveRoots struct {
	Axis string
}

func (b MuddyMangroveRoots) Encode() (string, BlockProperties) {
	return "minecraft:muddy_mangrove_roots", BlockProperties{
		"axis": b.Axis,
	}
}

func (b MuddyMangroveRoots) New(props BlockProperties) Block {
	return MuddyMangroveRoots{
		Axis: props["axis"],
	}
}