package main

import (
	"time"

	"github.com/wangbai/peeper/httpserv"
	"github.com/wangbai/peeper/monitor"
)

func main() {
	serv1 := monitor.NewService("http",
		[]monitor.Node{
			monitor.Node{Address: "baidu.com:80"},
			monitor.Node{Address: "sohu.com:80"},
		},
		3,
		2*time.Second,
	)

	serv2 := monitor.NewService("https",
		[]monitor.Node{
			monitor.Node{Address: "baidu.com:443"},
			monitor.Node{Address: "sohu.com:443"},
		},
		3,
		2*time.Second,
	)

	dis := monitor.NewDiscovery()
	dis.AddService(serv1)
	dis.AddService(serv2)
	dis.Start()

	monitor.RegisterDiscoveryLookupHandler(dis)
	httpserv.Start(11111)
}
