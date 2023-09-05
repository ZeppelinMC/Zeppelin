package packets

import (
	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/gui"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/util"
)

type server interface {
	PlayerlistRemove(players ...[16]byte)
	GlobalBroadcast(packet.Packet)
}

func Disconnect(conn *minecraft.Conn, srv server, logger logger.Logger) {
	uuid := util.ParseUUID(conn.Info.UUID)
	logger.Info("[%s] Player %s (%s) has left the server", conn.RemoteAddr().String(), conn.Info.Name, uuid)
	srv.PlayerlistRemove(conn.Info.UUID)
	gui.RemovePlayer(uuid)
}
