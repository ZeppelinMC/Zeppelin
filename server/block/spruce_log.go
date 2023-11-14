package block

import "github.com/dynamitemc/dynamite/server/world/chunk"

type SpruceLog struct {
	Stripped bool
	Axis     Axis
}

func (d SpruceLog) EncodedName() string {
	if d.Stripped {
		return "minecraft:stripped_spruce_log"
	}
	return "minecraft:spruce_log"
}

func (g SpruceLog) Properties() map[string]string {
	return map[string]string{
		"axis": g.Axis,
	}
}

func (SpruceLog) New(n string, p map[string]string) chunk.Block {
	return SpruceLog{
		Stripped: n == "minecraft:stripped_spruce_log",
		Axis:     p["axis"],
	}
}

func (SpruceLog) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 2,
	}
}
