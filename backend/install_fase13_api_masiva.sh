#!/data/data/com.termux/files/usr/bin/bash

set -e

echo "===== KORZADIVPN FASE 13 API VPNCORE ====="

mkdir -p handlers


cat > handlers/vpn_core_status.go <<'EOC'
package handlers

import (
	"encoding/json"
	"net/http"
)

func VPNCoreStatus(
	w http.ResponseWriter,
	r *http.Request,
){

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status":"online",
			"core":"wireguard",
			"nodes":"ready",
		},
	)
}
EOC


cat >> routes/routes.go <<'EOC'


// VPN CORE STATUS
// http://localhost:8080/api/vpn/core/status

EOC


gofmt -w handlers/vpn_core_status.go routes/routes.go

go build ./...

echo "===== FASE 13 API CREADA ====="

