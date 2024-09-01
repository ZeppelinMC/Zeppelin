package std

import (
	"math"
	"sync/atomic"

	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/util/log"
)

func chunkPos(x, z float64) (cx, cz int32) {
	return int32(math.Floor(x / 16)), int32(math.Floor(z / 16))
}

type Input struct {
	x, y, z    atomic.Uint64
	yaw, pitch atomic.Uint32
	onGround   atomic.Bool
}

func (i *Input) SetPosition(x, y, z float64) {
	i.x.Store(math.Float64bits(x))
	i.y.Store(math.Float64bits(y))
	i.z.Store(math.Float64bits(z))
}

func (i *Input) SetRotation(yaw, pitch float32) {
	i.yaw.Store(math.Float32bits(yaw))
	i.pitch.Store(math.Float32bits(pitch))
}

func (i *Input) SetOnGround(b bool) {
	i.onGround.Store(b)
}

func (i *Input) Position() (x, y, z float64) {
	return math.Float64frombits(i.x.Load()), math.Float64frombits(i.y.Load()), math.Float64frombits(i.z.Load())
}

func (i *Input) Rotation() (yaw, pitch float32) {
	return math.Float32frombits(i.yaw.Load()), math.Float32frombits(i.pitch.Load())
}

func (i *Input) OnGround() bool {
	return i.onGround.Load()
}

func (session *StandardSession) createTicker() {
	ticker := session.tick.New()
	go func() {
		for range ticker.C {
			session.processInput()
		}
	}()
}

func (session *StandardSession) processInput() {
	x, y, z, yaw, pitch, onGround := session.input()
	oldX, oldY, oldZ, oldYaw, oldPitch, oldOnGround := session.state()

	posC, rotC, onC := session.inputUpdated(x, y, z, yaw, pitch, onGround, oldX, oldY, oldZ, oldYaw, oldPitch, oldOnGround)
	if !posC && !rotC && !onC {
		return
	}

	if posC {
		oldChunkPosX, oldChunkPosZ := chunkPos(oldX, oldZ)
		newChunkPosX, newChunkPosZ := chunkPos(x, z)

		if oldChunkPosX != newChunkPosX || oldChunkPosZ != newChunkPosZ {
			session.WritePacket(&play.SetCenterChunk{ChunkX: newChunkPosX, ChunkZ: newChunkPosZ})
			session.ChunkLoadWorker.SendChunksRadius(newChunkPosX, newChunkPosZ)
		}

		distance := math.Sqrt((x-oldX)*(x-oldX) + (y-oldY)*(y-oldY) + (z-oldZ)*(z-oldZ))

		if distance > 100 {
			session.SynchronizePosition(oldX, oldY, oldZ, oldYaw, oldPitch)
			log.Infof("%s moved too quickly! (%f %f %f)\n", session.Username(), x-oldX, y-oldY, z-oldZ)
			return
		}
		defer session.player.SetPosition(x, y, z)
	}
	if rotC {
		defer session.player.SetRotation(yaw, pitch)
	}
	if onC {
		defer session.player.SetOnGround(onGround)
	}

	session.broadcast.BroadcastPlayerMovement(session, x, y, z, yaw, pitch)
}

func (session *StandardSession) inputUpdated(
	x, y, z float64, yaw, pitch float32, onGround bool,
	oldX, oldY, oldZ float64, oldYaw, oldPitch float32, oldOnGround bool,
) (posChanged, rotChanged, onGroundChanged bool) {

	return x != oldX || y != oldY || z != oldZ, yaw != oldYaw || pitch != oldPitch, onGround != oldOnGround
}

func (session *StandardSession) input() (x, y, z float64, yaw, pitch float32, onGround bool) {
	x, y, z = session.Input.Position()
	yaw, pitch = session.Input.Rotation()
	onGround = session.Input.OnGround()

	return
}

func (session *StandardSession) state() (oldX, oldY, oldZ float64, oldYaw, oldPitch float32, oldOnGround bool) {
	oldX, oldY, oldZ = session.player.Position()
	oldYaw, oldPitch = session.player.Rotation()
	oldOnGround = session.player.OnGround()

	return
}
