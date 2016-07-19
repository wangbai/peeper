package monitor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/wangbai/peeper/config"
)

func init() {
	config.Register("discovery", &DiscoveryConfig{})
}

type nodeItem struct {
	Address string            `json:"address"`
	Attr    map[string]string `json:"attr"`
}

type serviceItem struct {
	Name          string     `json:"name"`
	MaxFailureNum uint32     `json:"max_failure_num"`
	Interval      uint32     `json:"interval"`
	Timeout       uint32     `json:"timeout"`
	Nodes         []nodeItem `json:"nodes"`
}

type DiscoveryConfig []serviceItem

const configFile = "discovery.conf"

func (dc *DiscoveryConfig) ParseAndBuild(dir string) {
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
			nodes = append(nodes, Node{Address: n.Address, Attr: n.Attr})
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
	RegisterDiscoveryLookupHandler(dis)

    log.Printf("discovery app has been started");
}
