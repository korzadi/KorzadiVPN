package nodes

type Node struct {
	ID      int
	Name    string
	Country string
	IP      string
	Port    int
	Status  string
	Load    int
}

func NewNode(name string, ip string) Node {
	return Node{
		Name:   name,
		IP:     ip,
		Port:   51820,
		Status: "active",
	}
}
