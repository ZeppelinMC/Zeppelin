package command

import "github.com/zeppelinmc/zeppelin/net/packet/play"

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
		commandNodeIndex := int32(len(pk.Nodes))
		pk.Nodes[0].Children = append(pk.Nodes[0].Children, commandNodeIndex)
		pk.Nodes = append(pk.Nodes, cmd.Node.Node)

		if cmd.Namespace != "" {
			pk.Nodes[0].Children = append(pk.Nodes[0].Children, int32(len(pk.Nodes)))
			pk.Nodes = append(pk.Nodes, play.Node{
				Flags:        play.NodeLiteral | play.NodeRedirect,
				Name:         cmd.Namespace + ":" + cmd.Node.Name,
				RedirectNode: commandNodeIndex,
			})
		}

		for _, alias := range cmd.Aliases {
			pk.Nodes[0].Children = append(pk.Nodes[0].Children, int32(len(pk.Nodes)))
			pk.Nodes = append(pk.Nodes, play.Node{
				Flags:        play.NodeLiteral | play.NodeRedirect,
				Name:         alias,
				RedirectNode: commandNodeIndex,
			})

			if cmd.Namespace != "" {
				pk.Nodes[0].Children = append(pk.Nodes[0].Children, int32(len(pk.Nodes)))
				pk.Nodes = append(pk.Nodes, play.Node{
					Flags:        play.NodeLiteral | play.NodeRedirect,
					Name:         cmd.Namespace + ":" + alias,
					RedirectNode: commandNodeIndex,
				})
			}
		}

		for _, child := range cmd.Node.children {
			appendNode(commandNodeIndex, child, &pk.Nodes)
		}
	}

	return &pk
}

func appendNode(parentIndex int32, n Node, tgt *[]play.Node) {
	index := int32(len(*tgt))
	(*tgt)[parentIndex].Children = append((*tgt)[parentIndex].Children, index)
	(*tgt) = append((*tgt), n.Node)

	for _, child := range n.children {
		appendNode(index, child, tgt)
	}
}
