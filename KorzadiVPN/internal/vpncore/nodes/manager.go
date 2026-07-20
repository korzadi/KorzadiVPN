package nodes

type Manager struct {
	Nodes []Node
}

func NewManager() *Manager {
	return &Manager{
		Nodes: []Node{},
	}
}

func (m *Manager) AddNode(n Node) {
	m.Nodes = append(m.Nodes, n)
}

func (m *Manager) ActiveNodes() []Node {

	var result []Node

	for _, n := range m.Nodes {
		if n.Status == "active" {
			result = append(result, n)
		}
	}

	return result
}
