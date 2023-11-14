package tick

import (
	"time"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/block/pos"
	"github.com/dynamitemc/dynamite/server/entity"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/server/world/chunk"
	"github.com/google/uuid"
)

type Ticker struct {
	ticker  *time.Ticker
	server  *server.Server
	started bool
}

func New(srv *server.Server, tps int64) *Ticker {
	return &Ticker{ticker: time.NewTicker(time.Second / time.Duration(tps)), server: srv}
}

func (t *Ticker) Start() {
	if t.started {
		return
	}
	t.started = true
	go t.tickLoop()
}

func (t *Ticker) Restart(tps int64) {
	t.ticker = time.NewTicker(time.Second / time.Duration(tps))
}

func (t *Ticker) tickLoop() {
	var n uint
	for range t.ticker.C {
		t.tick(n)
		n++
	}
}

func (t *Ticker) tick(tick uint) {
	t.server.Entities.RangeNoLock(func(_ int32, e entity.Entity) bool {
		e.Tick(t.server, tick)
		return true
	})

	if tick%8 == 0 {
		blockTick(t.server, t.server.World.Overworld(), tick)
	}

	worldAge, dayTime := t.server.World.IncrementTime()
	t.server.Players.RangeNoLock(func(_ uuid.UUID, pl *player.Player) bool {
		if tick%8 == 0 {
			pl.SendChunks(pl.Dimension())
			//pl.UnloadChunks()
		}

		pl.SendPacket(&packet.UpdateTime{
			WorldAge:  worldAge,
			TimeOfDay: dayTime,
		})
		return true
	})
}

func blockTick(srv *server.Server, d *world.Dimension, tick uint) {
	for h, c := range d.Chunks() {
		cx, cz := h.Position()
		for cy, s := range c.Sections {
			for x := 0; x < 15; x++ {
				for y := 0; y < 15; y++ {
					for z := 0; z < 15; z++ {
						b := s.GetBlockAt(x, y, z)
						x1, y1, z1 := int64((int(cx)*16)+x), int64(((cy-4)*16)+y), int64((int(cz)*16)+z)

						if bl, ok := b.(interface {
							Tick(pos.BlockPosition, *world.Dimension, uint) chunk.Block
						}); ok {
							srv.SetBlock(d, x1, y1, z1, bl.Tick(pos.BlockPosition{x1, y1, z1}, d, tick), world.SetBlockReplace)
						}
					}
				}
			}
		}
	}
}
