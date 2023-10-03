package handlers

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func ClientSettings(state *player.Player, pk *packet.ClientSettings) {
	state.SetClientSettings(player.ClientInformation(*pk))
}
