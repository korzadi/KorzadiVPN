#!/data/data/com.termux/files/usr/bin/bash

set -e

echo "===== KORZADIVPN FASE 12 MASIVA ====="

mkdir -p vpncore/nodes
mkdir -p vpncore/wireguard
mkdir -p scripts


cat > vpncore/nodes/node.go <<'EOC'
package nodes

type Node struct {
	ID int
	Name string
	Country string
	IP string
	Port int
	Status string
	Load int
}

func NewNode(name string, ip string) Node {
	return Node{
		Name:name,
		IP:ip,
		Port:51820,
		Status:"active",
	}
}
EOC


cat > vpncore/nodes/manager.go <<'EOC'
package nodes

type Manager struct {
	Nodes []Node
}

func NewManager()*Manager{
	return &Manager{
		Nodes:[]Node{},
	}
}

func (m *Manager) AddNode(n Node){
	m.Nodes=append(m.Nodes,n)
}

func (m *Manager) ActiveNodes()[]Node{

	var result []Node

	for _,n:=range m.Nodes{
		if n.Status=="active"{
			result=append(result,n)
		}
	}

	return result
}
EOC


cat > vpncore/wireguard/server.go <<'EOC'
package wireguard

type Server struct {
	Name string
	IP string
	Port int
	PublicKey string
}

func NewServer(name string,ip string) Server{

	return Server{
		Name:name,
		IP:ip,
		Port:51820,
	}
}
EOC


gofmt -w vpncore

go build ./...

echo "===== FASE 12 INSTALADA ====="

