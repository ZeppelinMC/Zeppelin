package server

import (
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/log"
	"github.com/zeppelinmc/zeppelin/net/metadata"
	"github.com/zeppelinmc/zeppelin/net/packet/login"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/config"
	"github.com/zeppelinmc/zeppelin/server/entity"
	"github.com/zeppelinmc/zeppelin/server/player"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/world/region"
	"github.com/zeppelinmc/zeppelin/text"
)

var _ session.Session = (*Console)(nil)

// Console is an implementation of session that's the server
type Console struct {
	Server *Server
}

func (c *Console) Addr() net.Addr {
	return c.Server.listener.Addr()
}

func (c *Console) ClientName() string {
	return "zeppelin-console"
}

func (c *Console) Config() config.ServerConfig {
	return c.Server.cfg
}

func (c *Console) DespawnEntities(...int32) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) Dimension() *region.Dimension {
	return c.Server.World.Overworld
}

func (c *Console) Disconnect(text.TextComponent) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) EntityAnimation(int32, byte) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) EntityMetadata(int32, metadata.Metadata) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) IsSpawned(int32) bool {
	return false
}

func (c *Console) Player() *player.Player {
	return nil
}

func (c *Console) DisguisedChatMessage(msg text.TextComponent, session session.Session, chatType string) error {
	switch chatType {
	case "minecraft:chat":
		log.Printlnf("%s [Chat] <%s> %s", log.Time(), session.Username(), log.SprintText(msg))
	}
	return nil
}

func (c *Console) PlayerChatMessage(pk play.ChatMessage, session session.Session, chatType string, index int32, prev []play.PreviousMessage) error {
	switch chatType {
	case "minecraft:chat":
		log.Printlnf("%s [Chat] <%s> %s", log.Time(), session.Username(), pk.Message)
	}
	return nil
}

// will simply remove the players
func (c *Console) PlayerInfoRemove(uuids ...uuid.UUID) error {
	c.Server.Broadcast.RemoveUUIDs(text.TextComponent{Text: "Kicked from the server"}, uuids...)
	return nil
}

func (c *Console) PlayerInfoUpdate(pk *play.PlayerInfoUpdate) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) Properties() []login.Property {
	return nil
}

func (c *Console) SessionData() (play.PlayerSession, bool) {
	return play.PlayerSession{}, false
}

func (c *Console) SpawnEntity(entity.Entity) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) SpawnPlayer(session.Session) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) SynchronizePosition(float64, float64, float64, float32, float32) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) SystemMessage(msg text.TextComponent) error {
	log.Printlnf("%s [Chat] %s", log.Time(), log.SprintText(msg))
	return nil
}

func (c *Console) UUID() uuid.UUID {
	return uuid.Nil
}

func (c *Console) UpdateEntityPosition(entity.Entity, *play.UpdateEntityPosition) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) UpdateEntityPositionRotation(entity.Entity, *play.UpdateEntityPositionAndRotation) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) UpdateEntityRotation(entity.Entity, *play.UpdateEntityRotation) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) UpdateTime(int64, int64) error {
	return fmt.Errorf("unsupported function for console session")
}

func (c *Console) Textures() (login.Textures, error) {
	return login.Textures{}, fmt.Errorf("unsupported function for console session")
}

func (c *Console) Username() string {
	return "Console"
}
