package canvas

type Nodes []*Node

type Edges []*Edge

func (c *Canvas) FilterNodes() Nodes {
	return c.Nodes
}

func (s Nodes) ByType(t string) Nodes {
	var selection []*Node
	for _, node := range s {
		if node.Type == t {
			selection = append(selection, node)
		}
	}
	return selection
}
