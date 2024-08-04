package block

type MagentaWallBanner struct {
	Facing string
}

func (b MagentaWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:magenta_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b MagentaWallBanner) New(props BlockProperties) Block {
	return MagentaWallBanner{
		Facing: props["facing"],
	}
}