package block

import (
	"strconv"
)

type Lava struct {
	Level int
}

func (b Lava) Encode() (string, BlockProperties) {
	return "minecraft:lava", BlockProperties{
		"level": strconv.Itoa(b.Level),
	}
}

func (b Lava) New(props BlockProperties) Block {
	return Lava{
		Level: atoi(props["level"]),
	}
}