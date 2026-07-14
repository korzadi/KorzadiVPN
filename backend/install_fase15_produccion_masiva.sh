#!/data/data/com.termux/files/usr/bin/bash

set -e

echo "===== KORZADIVPN FASE 15 PRODUCCION ====="

mkdir -p deploy
mkdir -p deploy/scripts


cat > deploy/scripts/start.sh <<'EOC'
#!/data/data/com.termux/files/usr/bin/bash

cd "$(dirname "$0")/../../backend"

go run main.go
EOC


cat > deploy/scripts/update.sh <<'EOC'
#!/data/data/com.termux/files/usr/bin/bash

cd "$(dirname "$0")/../../backend"

git pull

go mod tidy

gofmt -w .

go build ./...
EOC


cat > deploy/README.txt <<'EOC'
KorzadiVPN Production Package

Incluye:

- Backend API
- Base SQLite
- VPN Core
- WireGuard
- Gestión de nodos
- Instaladores

Siguiente etapa:
Despliegue VPS + dominios + certificados SSL.
EOC


chmod +x deploy/scripts/*.sh

echo "===== CREANDO BACKUP FINAL ====="

mkdir -p ../backups/fase15

cp -r vpncore ../backups/fase15/
cp -r handlers ../backups/fase15/
cp -r database ../backups/fase15/


echo "===== COMPILACION FINAL ====="

go build ./...

echo "===== FASE 15 LISTA ====="

