package block

type JackOLantern struct {
	Facing string
}

func (b JackOLantern) Encode() (string, BlockProperties) {
	return "minecraft:jack_o_lantern", BlockProperties{
		"facing": b.Facing,
	}
}

func (b JackOLantern) New(props BlockProperties) Block {
	return JackOLantern{
		Facing: props["facing"],
	}
}