package handler

import (
	"github.com/dynamitemc/dynamite/server/enum"
	"github.com/dynamitemc/dynamite/server/player"
)

func ClientCommand(state *player.Player, action int32) {
	switch action {
	case enum.ClientCommandRespawn:
		{
			state.Respawn(state.Dimension())
		}
	}
}
