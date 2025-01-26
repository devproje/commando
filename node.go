package commando

import (
	"github.com/devproje/commando/types"
)

type Node struct {
	Name     string
	Desc     string
	Commando *Commando
	Handler  ArgHandler
	SubNodes []Node
	Opts     []types.OptionData
}

type ArgHandler func(n *Node) error

func (c *Commando) Register(name, desc string, handler ArgHandler, options ...types.OptionData) {
	node := Node{
		Name:     name,
		Desc:     desc,
		Handler:  handler,
		Commando: c,
		Opts:     options,
	}

	c.nodes = append(c.nodes, node)
}
