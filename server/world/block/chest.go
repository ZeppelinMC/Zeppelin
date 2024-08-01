package block

import (
	"fmt"
	"strconv"

	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/container"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/dimension"
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
	if props["type"] == "" {
		props["type"] = "single"
	}
	return Chest{
		Facing:      props["facing"],
		Type:        props["type"],
		Waterlogged: props["waterlogged"] == "true",
	}
}

func (g Chest) BlockEntity() chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:chest",
		Items: make(container.Container, 0),
	}
}

func (g Chest) Use(clicker session.Session, pk play.UseItemOn, dimension *dimension.Dimension) {
	w, ok := dimension.WindowManager.At(pk.BlockX, pk.BlockY, pk.BlockZ)
	if !ok {
		entity, ok := dimension.BlockEntity(pk.BlockX, pk.BlockY, pk.BlockZ)
		if !ok {
			fmt.Println("not found entity id", pk.BlockX, pk.BlockY, pk.BlockZ)
			return
		}
		if entity.Id != "minecraft:chest" {
			return
		}
		w = dimension.WindowManager.New("minecraft:generic_9x3", entity.Items, text.Sprint("Chest"))
		dimension.WindowManager.AddWindow([3]int32{pk.BlockX, pk.BlockY, pk.BlockZ}, w)
	}

	oldViewers := w.Viewers

	clicker.OpenWindow(w)
	clicker.Broadcast().BlockAction(pk.BlockX, pk.BlockY, pk.BlockZ, dimension.Name(), 1, w.Viewers)

	if oldViewers == 0 && w.Viewers > 0 { // chest was opened
		clicker.Broadcast().PlaySound(session.SoundEffect(
			"minecraft:block.chest.open", false, nil, play.SoundCategoryBlock, pk.BlockX, pk.BlockY, pk.BlockZ, 1, 1,
		), dimension.Name())
	}
}

var _ Block = (*Chest)(nil)
