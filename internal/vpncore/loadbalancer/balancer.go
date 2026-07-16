package loadbalancer

import (
	"korzadivpn/internal/vpncore/nodes/config"
	"korzadivpn/internal/vpncore/nodes/registry"
)

func SelectNode() (
	config.NodeConfig,
	bool,
) {

	nodes := registry.ListNodes()

	if len(nodes) == 0 {

		return config.NodeConfig{}, false
	}

	selected := nodes[0]

	for _, node := range nodes {

		if node.CurrentClients < selected.CurrentClients {

			selected = node
		}
	}

	return selected, true
}
