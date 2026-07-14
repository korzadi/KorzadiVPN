#!/data/data/com.termux/files/usr/bin/bash

set -e

echo "===== KORZADIVPN FASE 13 MASIVA ====="

mkdir -p vpncore/pool
mkdir -p vpncore/monitor


cat > vpncore/pool/ip_pool.go <<'EOC'
package pool

type IPPool struct {

	Available []string

	Used []string
}


func NewPool() *IPPool {

	return &IPPool{
		Available: []string{},
		Used: []string{},
	}
}


func (p *IPPool) Add(ip string){

	p.Available = append(
		p.Available,
		ip,
	)
}


func (p *IPPool) Allocate() string {

	if len(p.Available)==0{
		return ""
	}

	ip:=p.Available[0]

	p.Available=p.Available[1:]

	p.Used=append(
		p.Used,
		ip,
	)

	return ip
}
EOC


cat > vpncore/monitor/status.go <<'EOC'
package monitor

type ServerStatus struct {

	Server string

	Online bool

	Load int
}


func Check(server string) ServerStatus {

	return ServerStatus{

		Server: server,

		Online: true,

		Load:0,
	}
}
EOC


gofmt -w vpncore

go build ./...

echo "===== FASE 13 INSTALADA ====="

