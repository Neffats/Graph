package graph

type Node struct {
	value string
}

func NewNode(v string) *Node {
	return &Node{ value: v}
}

func (n *Node) String() string {
	return n.value
}

type Network struct {
	nodes []*Node
	edges map[Node][]*Node
}

func NewNetwork() *Network {
	return &Network{}
}

func (g *Network) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Network) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *Network) String() string {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	return s
}
