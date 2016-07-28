package monitor

import (
	"net"
	"time"
)

type Node struct {
	Address    string
	Status     bool
	FailureNum uint32
	Attr       map[string]string
}

type Service struct {
	Name          string
	Nodes         []Node
	MaxFailureNum uint32
	Interval      time.Duration
	Timeout       time.Duration
	nodeStatus    chan statusChan
}

type statusChan struct {
    index       int
    status      bool
}

func NewService(name string, nodes []Node, maxFailureNum uint32, interval time.Duration, timeout time.Duration) *Service {
	if len(nodes) < 1 {
		return nil
	}

	return &Service{
		Name:          name,
		Nodes:         nodes,
		MaxFailureNum: maxFailureNum,
		Interval:      interval,
		Timeout:       timeout,
		nodeStatus:    make(chan statusChan, len(nodes)),
	}
}

func (s *Service) Start() {
	//monitor all the node
	for index, node := range s.Nodes {
		go s.checkNodeStatus(index, node)
	}

	//update status
	go s.updateNodeStatus()
}

func (s *Service) updateNodeStatus() {
	for {
		st := <-s.nodeStatus
        status := st.status
        index := st.index        

		node := s.Nodes[index]
		if status {
			node.Status = status
			node.FailureNum = 0
		} else {
			node.FailureNum++
			if node.FailureNum >= s.MaxFailureNum {
				node.Status = status
			}
		}
		s.Nodes[index] = node
	}
}

func (s *Service) checkNodeStatus(index int, node Node) {
	for {
		var status bool
		err := tryConnect(node.Address, s.Timeout)

		if err != nil {
			status = false
		} else {
			status = true
		}

		s.nodeStatus <- statusChan {
			index: index,
			status:  status,
		}

		time.Sleep(s.Interval)
	}
}

func tryConnect(address string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err
	}

	defer conn.Close()
	return nil
}
