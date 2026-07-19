package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/internal/vpncore/nodes/config"
	"korzadivpn/internal/vpncore/nodes/registry"
)

func CreateVPNNode(
	w http.ResponseWriter,
	r *http.Request,
) {

	var node config.NodeConfig

	json.NewDecoder(
		r.Body,
	).Decode(&node)

	registry.AddNode(node)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		node,
	)
}

func ListVPNNodes(
	w http.ResponseWriter,
	r *http.Request,
) {

	nodes := registry.ListNodes()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		nodes,
	)
}
