package httpserv

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wangbai/peeper/monitor"
)

type node struct {
	Address string            `json:"address"`
	Status  bool              `json:"status"`
	Attr    map[string]string `json:"attr,omitempty"`
}

type services map[string][]node

type DiscoveryResource struct {
	dis *monitor.Discovery
}

func NewDiscoveryResource(d *monitor.Discovery) {
	if d == nil {
		log.Fatal("Discovery Service should not be empty")
	}

	return &DiscoveryResource{d}
}

func (d *DiscoveryResource) ContentType() string {
	return "application/json"
}

func (d *DiscoveryResource) HandleCommand(comm string) string {
	s := make(map[string][]node)

	s["rabbitmq"] = []node{
		node{":80", true, map[string]string{"type": "master"}},
		node{":81", true, nil},
	}

	s["mysql"] = []node{
		node{":80", true, nil},
		node{":81", true, nil},
	}

	j, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("%s", err)
	}

	return fmt.Sprintf("%s", j)
}
