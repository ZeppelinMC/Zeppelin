package block

import "github.com/dynamitemc/dynamite/server/world/chunk"

type OakLog struct {
	Stripped bool
	Axis     Axis
}

func (d OakLog) EncodedName() string {
	if d.Stripped {
		return "minecraft:stripped_oak_log"
	}
	return "minecraft:oak_log"
}

func (g OakLog) Properties() map[string]string {
	return map[string]string{
		"axis": g.Axis,
	}
}

func (OakLog) New(n string, p map[string]string) chunk.Block {
	return OakLog{
		Stripped: n == "minecraft:stripped_oak_log",
		Axis:     p["axis"],
	}
}

func (OakLog) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 2,
	}
}
