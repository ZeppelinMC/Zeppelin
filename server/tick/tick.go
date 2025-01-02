package tick

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	//"github.com/zeppelinmc/zeppelin/server/session"
)

// New creates a new tick manager with tps ticks per second
func New(tps int /*b *session.Broadcast*/) *TickManager {
	mgr := &TickManager{
		//b: b,
	}
	mgr.d.Store(int64(time.Second / time.Duration(tps)))

	return mgr
}

type TickManager struct {
	tickers []*time.Ticker

	mu sync.RWMutex
	d  atomic.Int64

	//b *session.Broadcast
}

func (mgr *TickManager) AddNew(f func()) {
	tick := mgr.New()
	go func() {
		for range tick.C {
			f()
		}
	}()
}

func (mgr *TickManager) SetFrequency(tps int) error {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	if tps == 0 {
		return fmt.Errorf("0 tps is not allowed. Use mgr.Freeze instead")
	}

	d := time.Second / time.Duration(tps)
	mgr.d.Store(int64(d))

	for _, ticker := range mgr.tickers {
		ticker.Reset(d)
	}

	/*mgr.b.Range(func(u uuid.UUID, s session.Session) bool {
		s.SetTickState(float32(tps), false)
		return true
	})*/

	return nil
}

func (mgr *TickManager) Freeze() {
	for _, ticker := range mgr.tickers {
		ticker.Stop()
	}
	/*mgr.b.Range(func(u uuid.UUID, s session.Session) bool {
		s.SetTickState(0, true)
		return true
	})*/
}

func (mgr *TickManager) Unfreeze() {
	freq := time.Duration(mgr.d.Load())
	for _, ticker := range mgr.tickers {
		ticker.Reset(freq)
	}
	//tps := 1 / (float32(freq) / float32(time.Second))
	/*mgr.b.Range(func(u uuid.UUID, s session.Session) bool {
		s.SetTickState(tps, false)
		return true
	})*/
}

func (mgr *TickManager) Add(ticker *time.Ticker) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.tickers = append(mgr.tickers, ticker)
}

func (mgr *TickManager) New() *time.Ticker {
	ticker := time.NewTicker(mgr.Frequency())
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.tickers = append(mgr.tickers, ticker)

	return ticker
}

func (mgr *TickManager) Count() int {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	return len(mgr.tickers)
}

func (mgr *TickManager) Remove(ticker *time.Ticker) (ok bool) {
	for _, t := range mgr.tickers {
		if t == ticker {
			return true
		}
	}
	return
}

func (mgr *TickManager) Frequency() time.Duration {
	return time.Duration(mgr.d.Load())
}
