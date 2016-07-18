package monitor

import (
	"log"
)

type Discovery struct {
	ss map[string]*Service
}

func NewDiscovery() *Discovery {
	return &Discovery{
		ss: make(map[string]*Service),
	}
}

func (d *Discovery) AddService(s *Service) bool {
	if s == nil {
		return false
	}

	_, ok := d.ss[s.Name]
	if ok {
		log.Fatal("Service:%s has existed", s.Name)
	}

	d.ss[s.Name] = s
	return true
}

func (d *Discovery) GetAllServices() []*Service {
	var services []*Service
	for _, s := range d.ss {
		services = append(services, s)
	}

	return services
}

func (d *Discovery) Start() {
	for _, s := range d.ss {
		s.Start()
	}
}
