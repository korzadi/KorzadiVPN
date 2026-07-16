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

	if len(a.IPs) == 0 {

		return ""
	}

	ip := a.IPs[0]

	a.IPs = a.IPs[1:]

	return ip
}
