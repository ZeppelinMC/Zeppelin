package command

import "github.com/zeppelinmc/zeppelin/net/packet/play"

// TODO add arguments
func (mgr *Manager) Encode() *play.Commands {
	if mgr.graph != nil {
		return mgr.graph
	}
	pk := play.Commands{
		Nodes: make([]play.Node, 1, len(mgr.commands)+1),
	}
	pk.Nodes[0] = play.Node{
		Children: make([]int32, 0, len(mgr.commands)),
	}

	for _, cmd := range mgr.commands {
		node := play.Node{
			Flags: play.NodeLiteral,
			Name:  cmd.Name,
		}

		i := len(pk.Nodes)
		pk.Nodes = append(pk.Nodes, node)
		pk.Nodes[0].Children = append(pk.Nodes[0].Children, int32(i))

		for _, alias := range cmd.Aliases {
			ai := len(pk.Nodes)
			n := play.Node{
				Name:         alias,
				Flags:        play.NodeLiteral | play.NodeRedirect,
				RedirectNode: int32(i),
			}
			pk.Nodes = append(pk.Nodes, n)
			pk.Nodes[0].Children = append(pk.Nodes[0].Children, int32(ai))
		}
	}

	return &pk
}
