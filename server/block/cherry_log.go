package block

import "github.com/dynamitemc/dynamite/server/world/chunk"

type CherryLog struct {
	Stripped bool
	Axis     Axis
}

func (d CherryLog) EncodedName() string {
	if d.Stripped {
		return "minecraft:stripped_cherry_log"
	}
	return "minecraft:cherry_log"
}

func (g CherryLog) Properties() map[string]string {
	return map[string]string{
		"axis": g.Axis,
	}
}

func (CherryLog) New(n string, p map[string]string) chunk.Block {
	return CherryLog{
		Stripped: n == "minecraft:stripped_cherry_log",
		Axis:     p["axis"],
	}
}

func (CherryLog) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 2,
	}
}
