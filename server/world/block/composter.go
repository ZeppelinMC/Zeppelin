package block

import (
	"strconv"
)

type Composter struct {
	Level int
}

func (b Composter) Encode() (string, BlockProperties) {
	return "minecraft:composter", BlockProperties{
		"level": strconv.Itoa(b.Level),
	}
}

func (b Composter) New(props BlockProperties) Block {
	return Composter{
		Level: atoi(props["level"]),
	}
}