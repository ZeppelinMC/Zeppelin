package handler

import (
	"github.com/dynamitemc/dynamite/server/enum"
	"github.com/dynamitemc/dynamite/server/player"
)

func PlayerAbilities(state *player.Player, flags byte) {
	state.SetFlying(flags == enum.PlayerAbilityFlying)
}
