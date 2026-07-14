package allocator

import (
	"korzadivpn/vpncore/loadbalancer"
)

type Allocation struct {
	User string

	Node string

	IP string
}

func AssignClient(
	email string,
	ip string,
) (
	Allocation,
	bool,
) {

	node, ok := loadbalancer.SelectNode()

	if !ok {

		return Allocation{}, false
	}

	return Allocation{

		User: email,

		Node: node.ID,

		IP: ip,
	}, true
}
