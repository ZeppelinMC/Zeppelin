package block

import (
	"strconv"
)

type Tnt struct {
	Unstable bool
}

func (b Tnt) Encode() (string, BlockProperties) {
	return "minecraft:tnt", BlockProperties{
		"unstable": strconv.FormatBool(b.Unstable),
	}
}

func (b Tnt) New(props BlockProperties) Block {
	return Tnt{
		Unstable: props["unstable"] != "false",
	}
}