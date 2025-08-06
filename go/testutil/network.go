package testutil

import (
	"net"
	"sync"
)

// GetFreePorts asks the kernel for free open ports that are ready to use.
func GetFreePorts(count int) ([]int, error) {
	var ports []int
	listeners := make([]*net.TCPListener, 0, count)

	defer func() {
		for _, l := range listeners {
			_ = l.Close()
		}
	}()

	for i := 0; i < count; i++ {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			return nil, err
		}

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			return nil, err
		}

		listeners = append(listeners, l)
		ports = append(ports, l.Addr().(*net.TCPAddr).Port)
	}

	return ports, nil
}

type freePorts struct {
	lock  sync.Mutex
	idx   int
	ports []int
}

func newFreePorts(ports []int) *freePorts {
	return &freePorts{
		idx:   0,
		ports: ports,
	}
}

func (p *freePorts) mustGetPort() int {
	defer p.lock.Unlock()
	p.lock.Lock()

	if p.idx == len(p.ports) {
		panic("no ports available")
	}

	port := p.ports[p.idx]
	p.idx++

	return port
}
