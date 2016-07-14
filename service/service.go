package service

import (
	"log"
	"net"
	"time"
)

type Node struct {
	Address    string
	Status     bool
	Interval   time.Duration
	Attributes map[string]string
}

type Service struct {
	Name       string
	Nodes      map[string]Node
	nodeStatus chan Node
}

func NewService(name string, nodes []Node) *Service {
	allNodes := make(map[string]Node)
	for _, node := range nodes {
		allNodes[node.Address] = node
	}

	return &Service{
		Name:       name,
		Nodes:      allNodes,
		nodeStatus: make(chan Node, len(nodes)),
	}
}

func (s *Service) StartMonitor() {
	//monitor all the node
	for _, node := range s.Nodes {
		go s.checkNodeStatus(node)
	}

	//update status
	go s.updateNodeStatus()
}

func (s *Service) updateNodeStatus() {
	for {
		status := <-s.nodeStatus

		node := s.Nodes[status.Address]
		node.Status = status.Status
		s.Nodes[status.Address] = node
	}
}

func (s *Service) checkNodeStatus(node Node) {
	log.Printf("Start to monitor Node:%s\n", node.Address)

	for {
		var status bool
		err := tryConnect(node.Address)

		if err != nil {
			status = false
			log.Printf("Node:%s can't be connect, %s", node.Address, err)
		} else {
			status = true
			log.Printf("Node:%s is fine", node.Address)
		}

		s.nodeStatus <- Node{
			Address: node.Address,
			Status:  status,
		}

		time.Sleep(node.Interval)
	}
}

func tryConnect(address string) error {
	timeout := 10 * time.Second
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err
	}

	defer conn.Close()
	return nil
}
