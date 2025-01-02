package block

import (
	"strconv"
)

type Water struct {
	Level int
}

func (b Water) Encode() (string, BlockProperties) {
	return "minecraft:water", BlockProperties{
		"level": strconv.Itoa(b.Level),
	}
}

func (b Water) New(props BlockProperties) Block {
	return Water{
		Level: atoi(props["level"]),
	}
}