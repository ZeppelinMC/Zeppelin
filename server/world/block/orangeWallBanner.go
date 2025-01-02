package block

type OrangeWallBanner struct {
	Facing string
}

func (b OrangeWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:orange_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b OrangeWallBanner) New(props BlockProperties) Block {
	return OrangeWallBanner{
		Facing: props["facing"],
	}
}