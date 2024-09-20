package server

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/metadata"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/protocol/properties"
	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server/entity"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
	"github.com/zeppelinmc/zeppelin/util/log"
)

var errConsoleUnsupportedFunc = errors.New("unsupported function for console session")

var _ session.DummySession = (*Console)(nil)

// Console is an implementation of session that's the server
type Console struct {
	Server *Server
}

func (c *Console) Config() properties.ServerProperties {
	return c.Server.cfg
}

func (c *Console) DespawnEntities(...int32) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) EntityAnimation(int32, byte) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) EntityMetadata(int32, metadata.Metadata) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) DisguisedChatMessage(msg text.TextComponent, session session.Session, chatType string) error {
	switch chatType {
	case "minecraft:chat":
		log.Printlnf("%s <%s> %s", log.Time(), session.Username(), log.SprintText(msg))
	}
	return nil
}

func (c *Console) PlayerChatMessage(pk play.ChatMessage, session session.Session, chatType string, index int32, prev []play.PreviousMessage) error {
	switch chatType {
	case "minecraft:chat":
		log.Printlnf("%s <%s> %s", log.Time(), session.Username(), pk.Message)
	}
	return nil
}

func (c *Console) PlayerInfoRemove(uuids ...uuid.UUID) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) PlayerInfoUpdate(pk *play.PlayerInfoUpdate) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) SpawnEntity(entity.Entity) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) SpawnPlayer(session.Session) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) SystemMessage(msg text.TextComponent) error {
	log.Printlnf("%s %s", log.Time(), log.SprintText(msg))
	return nil
}

func (c *Console) UpdateEntityPosition(entity.Entity, *play.UpdateEntityPosition) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) UpdateEntityPositionRotation(entity.Entity, *play.UpdateEntityPositionAndRotation) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) UpdateEntityRotation(entity.Entity, *play.UpdateEntityRotation) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) UpdateTime(int64, int64) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) BlockAction(*play.BlockAction) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) PlaySound(*play.SoundEffect) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) PlayEntitySound(*play.EntitySoundEffect) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) UpdateBlock(pos pos.BlockPosition, b section.Block) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) UpdateBlockEntity(pos pos.BlockPosition, be chunk.BlockEntity) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) DeleteMessage(id int32, sig [256]byte) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) DamageEvent(attacker, attacked session.Session, damageType string) error {
	return errConsoleUnsupportedFunc
}

func (c *Console) Broadcast() *session.Broadcast {
	return c.Server.World.Broadcast
}
