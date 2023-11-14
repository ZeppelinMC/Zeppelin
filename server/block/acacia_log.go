package block

import "github.com/dynamitemc/dynamite/server/world/chunk"

type AcaciaLog struct {
	Stripped bool
	Axis     Axis
}

func (d AcaciaLog) EncodedName() string {
	if d.Stripped {
		return "minecraft:stripped_acacia_log"
	}
	return "minecraft:acacia_log"
}

func (g AcaciaLog) Properties() map[string]string {
	return map[string]string{
		"axis": g.Axis,
	}
}

func (AcaciaLog) New(n string, p map[string]string) chunk.Block {
	return AcaciaLog{
		Stripped: n == "minecraft:stripped_acacia_log",
		Axis:     p["axis"],
	}
}

func (AcaciaLog) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 2,
	}
}
