package block

import (
	"strconv"
)

type SuspiciousGravel struct {
	Dusted int
}

func (b SuspiciousGravel) Encode() (string, BlockProperties) {
	return "minecraft:suspicious_gravel", BlockProperties{
		"dusted": strconv.Itoa(b.Dusted),
	}
}

func (b SuspiciousGravel) New(props BlockProperties) Block {
	return SuspiciousGravel{
		Dusted: atoi(props["dusted"]),
	}
}