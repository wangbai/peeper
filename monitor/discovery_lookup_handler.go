package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wangbai/peeper/httpserv"
)

// json decode helper for Node
type nodeOutput struct {
	Address string            `json:"address"`
	Status  bool              `json:"status"`
	Attr    map[string]string `json:"attr,omitempty"`
}

// json helper type for Service
type serviceOutput []nodeOutput

type discoveryLookupHandler struct {
	dis *Discovery
}

func registerDiscoveryLookupHandler(d *Discovery) {
	if d == nil {
		log.Fatal("Discovery Service should not be empty")
	}

	dh := &discoveryLookupHandler{dis: d}
	httpserv.Register("discovery", dh)
}

func (dh *discoveryLookupHandler) Handle(resp http.ResponseWriter, req *http.Request) {
	// compose json data
	jsonDis := make(map[string]serviceOutput)
	ss := dh.dis.GetAllServices()
	for _, s := range ss {
		var jsonServ []nodeOutput
		for _, n := range s.Nodes {
			jsonServ = append(jsonServ, nodeOutput {
				Address: n.Address,
				Status:  n.Status,
				Attr:    n.Attr,
			})
		}

		jsonDis[s.Name] = jsonServ
	}

	j, err := json.Marshal(jsonDis)
	if err != nil {
        resp.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resp, "%s", err)
	}

	// set header
	resp.Header().Set("Content-Type", "application/json")
    resp.WriteHeader(http.StatusOK)
	fmt.Fprintf(resp, "%s", j)
}
