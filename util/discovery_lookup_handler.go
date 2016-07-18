package monitor

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wangbai/peeper/httpserv"
)

// json decode helper for Node
type node struct {
	Address string            `json:"address"`
	Status  bool              `json:"status"`
	Attr    map[string]string `json:"attr,omitempty"`
}

// json helper type for Service
type service map[string][]node

// json helper type for Discovery
type discovery map[string]service

type discoveryHandler struct {
	dis *Discovery
}

func registerDiscoveryHandler(d *Discovery) {
	if d == nil 
		log.Fatal("Discovery Service should not be empty")
	}

    dh := &DiscoveryHandler{dis: d}
    httpserv.Register("discovery", dh)
}

func (dh *DiscoveryHandler) Handle(resp http.ResponseWriter, req *http.Request) {
    // set header
    resp.Header().Set("Content-Type", "application/json")

	s := make(map[string][]node)
    discover := d.dis;

    for _, si := range *discover {
        nodesInfo := si.Nodes
        var allNodes []node
        for _, ni := range nodesInfo {
            allNodes = append(allNodes, node {
                Address: ni.Address,
                Status: ni.Status,
                Attr: ni.Attr,
            });
        }

        s[si.Name] = allNodes;
    }

	j, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("%s", err)
	}

	return fmt.Sprintf("%s", j)
}
