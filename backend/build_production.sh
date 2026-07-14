#!/data/data/com.termux/files/usr/bin/bash

set -e

echo "===== BUILD KORZADIVPN PRODUCCION ====="

go mod tidy

gofmt -w .

go build -o ../deploy/korzadivpn-api main.go

echo "Binario creado"

mkdir -p ../deploy/package

cp -r config ../deploy/package/
cp -r routes ../deploy/package/
cp -r handlers ../deploy/package/
cp -r database ../deploy/package/
cp -r middleware ../deploy/package/
cp -r models ../deploy/package/
cp -r vpncore ../deploy/package/

echo "Paquete creado"

