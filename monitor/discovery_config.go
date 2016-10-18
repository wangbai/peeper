package monitor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/wangbai/peeper/server"
)

func init() {
	server.RegisterModule("discovery", &DiscoveryConfig{})
}

type nodeConfig struct {
	Address string            `json:"address"`
	Attr    map[string]string `json:"attr"`
}

type serviceConfig struct {
	Name          string       `json:"name"`
	MaxFailureNum uint32       `json:"max_failure_num"`
	Interval      uint32       `json:"interval"`
	Timeout       uint32       `json:"timeout"`
	Nodes         []nodeConfig `json:"nodes"`
}

type DiscoveryConfig []serviceConfig

const configFile = "discovery.conf"

func (dc *DiscoveryConfig) ParseAndLoad(dir string) {
	filePath := dir + "/" + configFile

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("When read ", filePath, " : ", err)
	}

	err = json.Unmarshal(file, dc)
	if err != nil {
		log.Fatal("When parse ", filePath, " : ", err)
	}

	dis := NewDiscovery()
	for _, serv := range *dc {
		var nodes []Node
		for _, n := range serv.Nodes {
			nodes = append(nodes, Node{Address: n.Address, Status: true, Attr: n.Attr})
		}

		dis.AddService(
			NewService(
				serv.Name,
				nodes,
				serv.MaxFailureNum,
				time.Duration(serv.Interval)*time.Second,
				time.Duration(serv.Timeout)*time.Second,
			),
		)
	}

	dis.Start()
	registerDiscoveryLookupHandler(dis)

	log.Printf("discovery module has been started")
}
