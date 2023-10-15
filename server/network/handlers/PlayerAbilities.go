package handlers

import (
	"github.com/dynamitemc/dynamite/server/player"
)

func PlayerAbilities(state *player.Player, flags byte) {
	state.SetFlying(flags == 0x02)
}
