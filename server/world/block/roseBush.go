package block

type RoseBush struct {
	Half string
}

func (b RoseBush) Encode() (string, BlockProperties) {
	return "minecraft:rose_bush", BlockProperties{
		"half": b.Half,
	}
}

func (b RoseBush) New(props BlockProperties) Block {
	return RoseBush{
		Half: props["half"],
	}
}