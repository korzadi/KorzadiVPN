package registry

import (
	"sync"

	"korzadivpn/internal/vpncore/nodes/config"
)

var (
	nodes = make(
		map[string]config.NodeConfig,
	)

	mutex sync.RWMutex
)

func AddNode(
	node config.NodeConfig,
) {

	mutex.Lock()

	defer mutex.Unlock()

	nodes[node.ID] = node
}

func GetNode(
	id string,
) (config.NodeConfig, bool) {

	mutex.RLock()

	defer mutex.RUnlock()

	node, ok := nodes[id]

	return node, ok
}

func ListNodes() []config.NodeConfig {

	mutex.RLock()

	defer mutex.RUnlock()

	result := []config.NodeConfig{}

	for _, node := range nodes {

		result = append(
			result,
			node,
		)
	}

	return result
}
