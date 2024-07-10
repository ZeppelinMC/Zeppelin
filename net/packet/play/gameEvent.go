package play

import "aether/net/io"

const (
	GameEventNoRespawnBlockAvailable = iota
	GameEventBeginRaining
	GameEventEndRaining
	GameEventChangeGamemode
	GameEventWinGame
	GameEventDemoEvent
	GameEventArrowHitPlayer
	GameEventRainLevelChange
	GameEventThunderLevelChange
	GameEventPlayPufferfishStingSound
	GameEventPlayElderGuardianMobAppearance
	GameEventEnableRespawnScreen
	GameEventLimitedCrafting
	GameEventStartWaitingChunks
)

//clientbound
const PacketIdGameEvent = 0x22

type GameEvent struct {
	Event byte
	Value float32
}

func (GameEvent) ID() int32 {
	return 0x22
}

func (g *GameEvent) Encode(w io.Writer) error {
	if err := w.Ubyte(g.Event); err != nil {
		return err
	}
	return w.Float(g.Value)
}

func (g *GameEvent) Decode(r io.Reader) error {
	if err := r.Ubyte(&g.Event); err != nil {
		return err
	}
	return r.Float(&g.Value)
}
