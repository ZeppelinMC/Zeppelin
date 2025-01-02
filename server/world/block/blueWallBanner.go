package block

type BlueWallBanner struct {
	Facing string
}

func (b BlueWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:blue_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b BlueWallBanner) New(props BlockProperties) Block {
	return BlueWallBanner{
		Facing: props["facing"],
	}
}