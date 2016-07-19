package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wangbai/peeper/httpserv"
)

// json decode helper for Node
type node struct {
	Address string            `json:"address"`
	Status  bool              `json:"status"`
	Attr    map[string]string `json:"attr,omitempty"`
}

// json helper type for Service
type service []node

// json helper type for Discovery
type discovery map[string]service

type discoveryLookupHandler struct {
	dis *Discovery
}

func RegisterDiscoveryLookupHandler(d *Discovery) {
	if d == nil {
		log.Fatal("Discovery Service should not be empty")
	}

	dh := &discoveryLookupHandler{dis: d}
	httpserv.Register("discovery", dh)
}

func (dh *discoveryLookupHandler) Handle(resp http.ResponseWriter, req *http.Request) {
	// set header
	resp.Header().Set("Content-Type", "application/json")

	// compose json data
	jsonDis := make(map[string]service)
	ss := dh.dis.GetAllServices()
	for _, s := range ss {
		var jsonServ []node
		for _, n := range s.Nodes {
			jsonServ = append(jsonServ, node{
				Address: n.Address,
				Status:  n.Status,
				Attr:    n.Attr,
			})
		}

		jsonDis[s.Name] = jsonServ
	}

	j, err := json.Marshal(jsonDis)
	if err != nil {
		fmt.Fprintf(resp, "%s", err)
	}

	fmt.Fprintf(resp, "%s", j)
}
