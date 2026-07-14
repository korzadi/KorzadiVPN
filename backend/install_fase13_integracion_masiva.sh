#!/data/data/com.termux/files/usr/bin/bash

set -e

echo "===== KORZADIVPN FASE 13 INTEGRACION ====="

mkdir -p vpncore/allocator


cat > vpncore/allocator/allocator.go <<'EOC'
package allocator

import (
	"sync"
)

type Allocator struct {

	mu sync.Mutex

	IPs []string
}


func NewAllocator() *Allocator {

	return &Allocator{
		IPs: []string{
			"10.8.0.2",
			"10.8.0.3",
			"10.8.0.4",
			"10.8.0.5",
			"10.8.0.6",
		},
	}
}


func (a *Allocator) Next() string {

	a.mu.Lock()

	defer a.mu.Unlock()

	if len(a.IPs)==0 {

		return ""
	}

	ip:=a.IPs[0]

	a.IPs=a.IPs[1:]

	return ip
}
EOC


cat > vpncore/allocator/server.go <<'EOC'
package allocator

type Assignment struct {

	Email string

	IP string

	Server string
}


func Assign(
	email string,
	server string,
	ip string,
) Assignment {

	return Assignment{

		Email:email,

		IP:ip,

		Server:server,
	}
}
EOC


gofmt -w vpncore

go build ./...

echo "===== FASE 13 COMPLETA ====="

