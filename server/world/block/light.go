package block

import (
	"strconv"
)

type Light struct {
	Waterlogged bool
	Level int
}

func (b Light) Encode() (string, BlockProperties) {
	return "minecraft:light", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"level": strconv.Itoa(b.Level),
	}
}

func (b Light) New(props BlockProperties) Block {
	return Light{
		Waterlogged: props["waterlogged"] != "false",
		Level: atoi(props["level"]),
	}
}