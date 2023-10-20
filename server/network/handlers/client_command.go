package handlers

import (
	"github.com/dynamitemc/dynamite/server/player"
)

func ClientCommand(controller Controller, state *player.Player, action int32) {
	switch action {
	case 0:
		{
			controller.Respawn(state.Dimension())
		}
	}
}
