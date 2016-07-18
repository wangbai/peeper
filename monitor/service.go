package monitor

import (
    "net"
	"time"
)

type Node struct {
	Address    string
	Status     bool
	FailureNum uint32
	Attr map[string]string
}

type Service struct {
	Name          string
	Nodes         map[string]Node
	MaxFailureNum uint32
	Interval      time.Duration
	nodeStatus    chan Node
}

func NewService(name string, nodes []Node, maxFailureNum uint32, interval time.Duration) *Service {
	if len(nodes) < 1 {
		return nil
	}

	allNodes := make(map[string]Node)
	for _, node := range nodes {
		allNodes[node.Address] = node
	}

	return &Service{
		Name:          name,
		Nodes:         allNodes,
		MaxFailureNum: maxFailureNum,
		Interval:      interval,
		nodeStatus:    make(chan Node, len(nodes)),
	}
}

func (s *Service) Start() {
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
		if status.Status {
			node.Status = status.Status
			node.FailureNum = 0
		} else {
			node.FailureNum++
			if node.FailureNum >= s.MaxFailureNum {
				node.Status = status.Status
			}
		}
		s.Nodes[status.Address] = node
	}
}

func (s *Service) checkNodeStatus(node Node) {
	for {
		var status bool
		err := tryConnect(node.Address)

		if err != nil {
			status = false
		} else {
			status = true
		}

		s.nodeStatus <- Node{
			Address: node.Address,
			Status:  status,
		}

		time.Sleep(s.Interval)
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
