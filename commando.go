package commando

import (
	"fmt"
	"github.com/devproje/commando/types"
)

type Commando struct {
	args  []string
	nodes []Node
}

type Node struct {
	Name     string
	Desc     string
	Commando *Commando
	Handler  ArgHandler
	SubNodes []Node
	Opts     []types.OptionData
}

type ArgHandler func(n *Node) error

func NewCommando(args []string) *Commando {
	return &Commando{nodes: make([]Node, 0), args: args}
}

func (c *Commando) Args() []string {
	return c.args
}

func (c *Commando) Nodes() []Node {
	return c.nodes
}

func (c *Commando) Execute() error {
	if len(c.args) == 0 {
		return fmt.Errorf("no command specified")
	}

	cmd := c.args[0]
	var node *Node

	for _, n := range c.nodes {
		if n.Name == cmd {
			node = &n
		}
	}

	if node == nil {
		return fmt.Errorf("no command name: %s", cmd)
	}

	var err error
	handler := node.Handler
	if handler == nil {
		if len(node.SubNodes) > 0 {
			if len(c.args) == 1 {
				return fmt.Errorf("no command specified: %s", c.args[0])
			}

			cmd = c.args[1]
			for _, n := range node.SubNodes {
				if n.Name == cmd {
					node = &n
					break
				}
			}

			if node == nil {
				return fmt.Errorf("no command name: %s", cmd)
			}

			handler = node.Handler
			err = handler(node)
			if err != nil {
				return err
			}

			return nil
		}

		return fmt.Errorf("no command handler: %s", cmd)
	}

	err = handler(node)
	if err != nil {
		return err
	}

	return nil
}

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
