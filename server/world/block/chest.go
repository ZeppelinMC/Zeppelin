package block

import (
	"strconv"

	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/world/dimension"
	"github.com/zeppelinmc/zeppelin/server/world/dimension/window"
	"github.com/zeppelinmc/zeppelin/text"
)

type Chest struct {
	Facing      string
	Type        string
	Waterlogged bool
}

func (g Chest) Encode() (string, BlockProperties) {
	return "minecraft:chest", BlockProperties{
		"facing":      g.Facing,
		"type":        g.Type,
		"waterlogged": strconv.FormatBool(g.Waterlogged),
	}
}

func (g Chest) New(props BlockProperties) Block {
	return Chest{
		Facing:      props["facing"],
		Type:        props["type"],
		Waterlogged: props["waterlogged"] == "true",
	}
}

func (g Chest) Use(clicker session.Session, pk play.UseItemOn, dimension *dimension.Dimension) {
	w, ok := dimension.WindowManager.At(pk.BlockX, pk.BlockY, pk.BlockZ)
	if !ok {
		entity, ok := dimension.BlockEntity(pk.BlockX, pk.BlockY, pk.BlockZ)
		if !ok {
			return
		}
		if entity.Id != "minecraft:chest" {
			return
		}
		w = window.New("minecraft:generic_9x3", entity.Items, text.Sprint("Chest"))
		dimension.WindowManager.AddWindow([3]int32{pk.BlockX, pk.BlockY, pk.BlockZ}, w)
	}
	clicker.OpenWindow(*w)
	w.Viewers++
	clicker.Broadcast().BlockAction(pk.BlockX, pk.BlockY, pk.BlockZ, dimension.Name(), 1, w.Viewers)
}

var _ Block = (*Chest)(nil)
