#!/data/data/com.termux/files/usr/bin/bash

set -e

echo "===== KORZADIVPN FASE 14 BASE APP ====="

mkdir -p mobile/api
mkdir -p mobile/models


cat > mobile/models/config.go <<'EOC'
package models

type ServerConfig struct {

	API string

	VPN string

	Version string
}
EOC


cat > mobile/api/client.go <<'EOC'
package api

type Client struct {

	BaseURL string
}


func NewClient(
	url string,
) Client {

	return Client{

		BaseURL:url,
	}
}
EOC


cat > mobile/README.txt <<'EOC'
KorzadiVPN Mobile

Módulos preparados:

- Login
- Dashboard
- Servidores
- Conexión VPN
- Perfil WireGuard

Pendiente:
Implementación Android nativa.
EOC


echo "===== MOBILE BASE CREADA ====="

