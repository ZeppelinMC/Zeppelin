package block

type YellowWallBanner struct {
	Facing string
}

func (b YellowWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:yellow_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b YellowWallBanner) New(props BlockProperties) Block {
	return YellowWallBanner{
		Facing: props["facing"],
	}
}