package pool

type IPPool struct {
	Available []string

	Used []string
}

func NewPool() *IPPool {

	return &IPPool{
		Available: []string{},
		Used:      []string{},
	}
}

func (p *IPPool) Add(ip string) {

	p.Available = append(
		p.Available,
		ip,
	)
}

func (p *IPPool) Allocate() string {

	if len(p.Available) == 0 {
		return ""
	}

	ip := p.Available[0]

	p.Available = p.Available[1:]

	p.Used = append(
		p.Used,
		ip,
	)

	return ip
}
