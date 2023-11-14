package block

import (
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type Axis = string

const (
	AxisX Axis = "x"
	AxisY Axis = "y"
	AxisZ Axis = "z"
)

type DarkOakLog struct {
	Stripped bool
	Axis     Axis
}

func (d DarkOakLog) EncodedName() string {
	if d.Stripped {
		return "minecraft:stripped_dark_oak_log"
	}
	return "minecraft:dark_oak_log"
}

func (g DarkOakLog) Properties() map[string]string {
	return map[string]string{
		"axis": g.Axis,
	}
}

func (DarkOakLog) New(n string, p map[string]string) chunk.Block {
	return DarkOakLog{
		Stripped: n == "minecraft:stripped_dark_oak_log",
		Axis:     p["axis"],
	}
}

func (DarkOakLog) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 2,
	}
}
