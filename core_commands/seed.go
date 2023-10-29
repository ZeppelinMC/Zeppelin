package core_commands

import (
	"fmt"

	"github.com/aimjel/minecraft/chat"
	"github.com/dynamitemc/dynamite/server/commands"
)

var seed_cmd = &commands.Command{
	Name:                "seed",
	RequiredPermissions: []string{"server.command.seed"},
	Execute: func(ctx commands.CommandContext) {
		server := getServer(ctx.Executor)
		seed := server.World.Seed()
		ctx.Reply(server.Translate("commands.seed.success", chat.NewMessage(fmt.Sprint(seed))).WithCopyToClipboardClickEvent(fmt.Sprint(seed)).WithShowTextHoverEvent(chat.NewMessage("Click to Copy to clipboard")))
	},
}
