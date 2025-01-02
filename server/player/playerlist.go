package player

import (
	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/util/atomic"
	"unsafe"
)

func NewPlayerList(cap uintptr) PlayerList {
	return PlayerList{
		playerList: atomic.Make(0, cap, atomic.Pointer),
	}
}

type PlayerList struct {
	playerList atomic.Slice
}

func (list *PlayerList) NumPlayers() int {
	return int(list.playerList.Len())
}

func (list *PlayerList) Player(uuid uuid.UUID) *Player {
	for i := uintptr(0); i < list.playerList.Len(); i++ {
		player := list.playerAtIndex(i)

		if player.UUID() == uuid {
			return player
		}
	}
	return nil
}

func (list *PlayerList) PlayerByUsername(name string) *Player {
	for i := uintptr(0); i < list.playerList.Len(); i++ {
		player := list.playerAtIndex(i)

		if player.Username() == name {
			return player
		}
	}
	return nil
}

func (list *PlayerList) playerAtIndex(index uintptr) *Player {
	return *(**Player)(list.playerList.Element(index))
}

func (list *PlayerList) AddPlayer(player *Player) {
	list.addPlayerToSlice(player)

	update := &play.PlayerInfoUpdate{
		Actions: play.ActionAddPlayer | play.ActionUpdateListed, //todo: add the rest
		Players: make(map[uuid.UUID]play.PlayerAction, list.NumPlayers()),
	}

	for i := uintptr(0); i < list.playerList.Len(); i++ {
		player := list.playerAtIndex(i)

		update.Players[player.UUID()] = play.PlayerAction{
			Name:       player.Username(),
			Listed:     !player.Unlisted,
			Properties: player.Properties(),
		}
	}

	for i := uintptr(0); i < list.playerList.Len(); i++ {
		player := list.playerAtIndex(i)
		player.writeOrKill(update)
	}
}

func (list *PlayerList) addPlayerToSlice(player *Player) {
	pointerOfPlayer := unsafe.Pointer(player)
	list.playerList.Append(unsafe.Pointer(&pointerOfPlayer))
}
