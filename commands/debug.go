package commands

import (
	"math"

	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server"
	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/util"
)

var debug = command.Command{
	Node:      command.NewLiteral("debug"),
	Aliases:   []string{"f3"},
	Namespace: "zeppelin",
	Callback: func(ccc command.CommandCallContext) {
		s, ok := ccc.Executor.(session.Session)
		if !ok {
			ccc.Executor.SystemMessage(text.TextComponent{
				Text:  "This command should be used by a player.",
				Color: "red",
			})
			return
		}
		player := s.Player()
		if player == nil {
			ccc.Executor.SystemMessage(text.TextComponent{
				Text:  "This command should be used by a player.",
				Color: "red",
			})
			return
		}
		x, y, z := player.Position()
		chunkX, chunkY, chunkZ := int32(x)>>4, int32(y)>>4, int32(z)>>4
		xb, yb, zb := int32(math.Floor(x)), int32(math.Floor(y)), int32(math.Floor(z))
		rx, rz := chunkX>>5, chunkZ>>5

		yaw, pitch := player.Rotation()

		dimension := ccc.Server.(*server.Server).World.Dimension(player.Dimension())

		c, err := dimension.GetChunk(chunkX, chunkZ)
		if err != nil {
			ccc.Executor.SystemMessage(text.TextComponent{
				Text:  "Unrendered chunk",
				Color: "red",
			})
			return
		}
		onBlock, _ := c.Block(xb&0x0f, yb-1, zb&0x0f)
		sky, _ := c.SkyLightLevel(xb&0x0f, yb, zb&0x0f)
		block, _ := c.BlockLightLevel(xb&0x0f, yb, zb&0x0f)

		name, props := onBlock.Encode()

		ccc.Executor.SystemMessage(text.Unmarshalf(
			ccc.Executor.Config().ChatFormatter.Rune(),
			"XYZ: %.03f / %.05f / %.03f\nBlock: %d %d %d [%d %d %d]\nChunk: %d %d %d [%d %d in r.%d.%d.mca]\nStanding on: %s [%v]\nFacing: (%.01f / %.01f)\nClient Light: %d (%d sky, %d block)\n\nYou are using: %s",
			x, y, z,
			xb, yb, zb,
			xb&0xf, yb&0xf, zb&0xf,
			chunkX, chunkY, chunkZ,
			chunkX&31, chunkZ&31,
			rx, rz,
			name, props,
			util.NormalizeYaw(yaw), pitch,
			sky+block, sky, block,
			s.ClientName(),
		))
	},
}
