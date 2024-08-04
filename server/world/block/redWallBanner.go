package block

type RedWallBanner struct {
	Facing string
}

func (b RedWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:red_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b RedWallBanner) New(props BlockProperties) Block {
	return RedWallBanner{
		Facing: props["facing"],
	}
}