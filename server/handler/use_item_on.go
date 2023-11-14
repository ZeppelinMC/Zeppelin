package handler

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

func UseItemOn(state *player.Player, pk *packet.UseItemOnServer, f func(d *world.Dimension, x, y, z int64, b chunk.Block, typ world.SetBlockHandling)) {
	x, y, z := world.ParsePosition(int64(pk.Location))

	if b := state.Dimension().Block(x, y, z); b != nil && b.EncodedName() != "minecraft:air" {

		// todo check for snow/ flowers etc
		//return
	}
	i, ok := state.Inventory.Slot(int8(state.SelectedSlot()))
	if !ok {
		return
	}

	b := chunk.DefaultBlock(i.Id)
	if b == nil || b.EncodedName() == "minecraft:air" {
		return
	}

	f(state.Dimension(), x, y, z, b, world.SetBlockReplace)
}
