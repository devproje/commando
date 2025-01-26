package commando

type Node struct {
	Name     string
	Desc     string
	Commando *Commando
	Handler  ArgHandler
	SubNodes []Node
}

type ArgHandler func(n *Node) error

func (c *Commando) Register(name, desc string, handler ArgHandler) {
	node := Node{
		Name:     name,
		Desc:     desc,
		Handler:  handler,
		Commando: c,
	}

	c.nodes = append(c.nodes, node)
}
