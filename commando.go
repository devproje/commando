package commando

import (
	"fmt"
)

type Commando struct {
	args  []string
	nodes []Node
	Must  bool
	Help  bool
}

func NewCommando(args []string) *Commando {
	return &Commando{nodes: make([]Node, 0), args: args, Must: true, Help: true}
}

func (c *Commando) SetMust(value bool) {
	c.Must = value
}

func (c *Commando) SetHelp(value bool) {
	c.Help = value
}

func (c *Commando) Args() []string {
	return c.args
}

func (c *Commando) Nodes() []Node {
	return c.nodes
}

func (c *Commando) Execute() error {
	if len(c.args) == 0 {
		if !c.Must {
			return nil
		}

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
