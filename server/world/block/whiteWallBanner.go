package block

type WhiteWallBanner struct {
	Facing string
}

func (b WhiteWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:white_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b WhiteWallBanner) New(props BlockProperties) Block {
	return WhiteWallBanner{
		Facing: props["facing"],
	}
}