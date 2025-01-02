package block

type LightBlueWallBanner struct {
	Facing string
}

func (b LightBlueWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b LightBlueWallBanner) New(props BlockProperties) Block {
	return LightBlueWallBanner{
		Facing: props["facing"],
	}
}