package command

import (
	"strings"
	"sync"

	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/text"
)

type Manager struct {
	commands []Command
	mu       sync.RWMutex
	srv      any

	graph *play.Commands
}

func NewManager(srv any, cmds ...Command) *Manager {
	return &Manager{commands: cmds, srv: srv}
}

func (mgr *Manager) Register(cmds ...Command) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.commands = append(mgr.commands, cmds...)
	mgr.graph = nil
}

func (mgr *Manager) Call(command string, caller session.Session) {
	arguments := strings.Split(command, " ")
	if len(arguments) == 0 {
		caller.SystemMessage(
			text.Unmarshal(
				"&cInvalid command", '&',
			),
		)
		return
	}
	cmd := mgr.findCommand(arguments[0])
	if cmd == nil {
		caller.SystemMessage(
			text.Unmarshalf(
				'&',
				"&cUnknown command %s",
				command,
			),
		)
		return
	}
	ctx := CommandCallContext{
		Command:  *cmd,
		Executor: caller,
		Server:   mgr.srv,
	}
	if len(arguments) > 1 {
		ctx.Arguments = arguments[1:]
	}
	cmd.Callback(ctx)
}

func (mgr *Manager) findCommand(name string) *Command {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	var namespace string
	if i := strings.Index(name, ":"); i != -1 && i != len(name) {
		namespace = name[:i]
		name = name[i+1:]
	}
	for _, cmd := range mgr.commands {
		if namespace != "" && cmd.Namespace != namespace {
			continue
		}
		if cmd.Node.Name == name {
			return &cmd
		}
		for _, alias := range cmd.Aliases {
			if alias == name {
				return &cmd
			}
		}
	}
	return nil
}
