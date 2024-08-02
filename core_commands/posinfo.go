package core_commands

import (
	"math"

	"github.com/zeppelinmc/zeppelin/server"
	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/text"
)

var posinfo = command.Command{
	Name: "posinfo",
	Callback: func(ccc command.CommandCallContext) {
		player := ccc.Executor.Player()
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

		name, props := onBlock.Encode()

		ccc.Executor.SystemMessage(text.Unmarshalf(
			ccc.Executor.Config().Chat.Formatter.Rune(),
			"XYZ: %.03f / %.05f / %.03f\nBlock: %d %d %d [%d %d %d]\nChunk: %d %d %d [%d %d in r.%d.%d.mca]\nStanding on: %s [%v]",
			x, y, z,
			xb, yb, zb,
			xb&0xf, yb&0xf, zb&0xf,
			chunkX, chunkY, chunkZ,
			chunkX&31, chunkZ&31,
			rx, rz,
			name, props,
		))
	},
}
