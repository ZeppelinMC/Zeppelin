package handler

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/enum"
	"github.com/dynamitemc/dynamite/server/player"
)

var (
	st       = enum.PoseStanding
	sw       = enum.PoseSwimming
	sn       = enum.PoseSneaking
	dg       = enum.PoseDigging
	sp  byte = 0x08
	nsp byte
)

func PlayerCommand(state *player.Player, action int32) {
	var pk = &packet.SetEntityMetadata{
		EntityID: state.EntityID(),
		Pose:     &dg,
	}
	switch action {
	case enum.PlayerCommandStartSneaking: // start sneaking / swimming
		{
			b := state.OnBlock()
			if b.EncodedName() == "minecraft:water" {
				pk.Pose = &sw
			} else {
				pk.Pose = &sn
			}
		}
	case enum.PlayerCommandStopSneaking: // stop sneaking / swimming
		{
			pk.Pose = &st
		}
	case enum.PlayerCommandStartSprinting: // sprint
		{
			pk.Data = &sp
		}
	case enum.PlayerCommandStopSprinting: // stop sprint
		{
			pk.Data = &nsp
		}
	}
	state.BroadcastMetadataInArea(pk)
}
