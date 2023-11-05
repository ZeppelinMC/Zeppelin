package handlers

import (
	"github.com/dynamitemc/dynamite/server/enum"
	"github.com/dynamitemc/dynamite/server/player"
)

func ClientCommand(controller Controller, state *player.Player, action int32) {
	switch action {
	case enum.ClientCommandRespawn:
		{
			controller.Respawn(state.Dimension())
		}
	}
}
