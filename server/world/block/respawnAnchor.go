package block

import (
	"strconv"
)

type RespawnAnchor struct {
	Charges int
}

func (b RespawnAnchor) Encode() (string, BlockProperties) {
	return "minecraft:respawn_anchor", BlockProperties{
		"charges": strconv.Itoa(b.Charges),
	}
}

func (b RespawnAnchor) New(props BlockProperties) Block {
	return RespawnAnchor{
		Charges: atoi(props["charges"]),
	}
}