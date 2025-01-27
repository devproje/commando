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

func (c *Commando) Root(name, desc string, handler ArgHandler, options ...types.OptionData) {
	c.nodes = append(c.nodes, Node{
		Name:     name,
		Desc:     desc,
		Commando: c,
		Handler:  handler,
		SubNodes: nil,
		Opts:     options,
	})
}

func (c *Commando) ComplexRoot(name string, desc string, subNodes []Node, options ...types.OptionData) {
	c.nodes = append(c.nodes, Node{
		Name:     name,
		Desc:     desc,
		Commando: c,
		SubNodes: subNodes,
		Opts:     options,
	})
}

func (c *Commando) Then(name string, desc string, handler ArgHandler, options ...types.OptionData) Node {
	return Node{
		Name:     name,
		Desc:     desc,
		Commando: c,
		Handler:  handler,
		SubNodes: nil,
		Opts:     options,
	}
}

func (n *Node) MustGetOpt(name string) *types.OptionData {
	for _, opt := range n.Opts {
		if opt.Name == name {
			return &opt
		}
	}

	return nil
}
