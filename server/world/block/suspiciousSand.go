package block

import (
	"strconv"
)

type SuspiciousSand struct {
	Dusted int
}

func (b SuspiciousSand) Encode() (string, BlockProperties) {
	return "minecraft:suspicious_sand", BlockProperties{
		"dusted": strconv.Itoa(b.Dusted),
	}
}

func (b SuspiciousSand) New(props BlockProperties) Block {
	return SuspiciousSand{
		Dusted: atoi(props["dusted"]),
	}
}