package block

import "github.com/dynamitemc/dynamite/server/world/chunk"

type BirchLog struct {
	Stripped bool
	Axis     Axis
}

func (d BirchLog) EncodedName() string {
	if d.Stripped {
		return "minecraft:stripped_birch_log"
	}
	return "minecraft:birch_log"
}

func (g BirchLog) Properties() map[string]string {
	return map[string]string{
		"axis": g.Axis,
	}
}

func (BirchLog) New(n string, p map[string]string) chunk.Block {
	return BirchLog{
		Stripped: n == "minecraft:stripped_birch_log",
		Axis:     p["axis"],
	}
}

func (BirchLog) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 2,
	}
}
