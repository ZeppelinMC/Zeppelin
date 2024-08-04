package block

type PurpleWallBanner struct {
	Facing string
}

func (b PurpleWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:purple_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b PurpleWallBanner) New(props BlockProperties) Block {
	return PurpleWallBanner{
		Facing: props["facing"],
	}
}