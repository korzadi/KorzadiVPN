#!/data/data/com.termux/files/usr/bin/bash

set -e

echo "===== KORZADIVPN FASE 14 API CLIENTE ====="

mkdir -p handlers
mkdir -p models


cat > models/app_config.go <<'EOC'
package models

type AppConfig struct {

	Version string `json:"version"`

	API string `json:"api"`

	VPNCore string `json:"vpn_core"`

	Status string `json:"status"`
}
EOC


cat > handlers/app_config.go <<'EOC'
package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/models"
)


func AppConfig(
	w http.ResponseWriter,
	r *http.Request,
){

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		models.AppConfig{

			Version:"1.0",

			API:"KorzadiVPN",

			VPNCore:"WireGuard",

			Status:"ready",
		},
	)
}
EOC


python3 - <<'PY'
from pathlib import Path

p=Path("routes/routes.go")

s=p.read_text()

if '"/api/app/config"' not in s:

    s=s.replace(
        'func RegisterRoutes() {',
        'func RegisterRoutes() {\\n\\n\\thttp.HandleFunc("/api/app/config", handlers.AppConfig)'
    )

p.write_text(s)
PY


gofmt -w models/app_config.go handlers/app_config.go routes/routes.go

go build ./...

echo "===== FASE 14 API CLIENTE LISTA ====="

