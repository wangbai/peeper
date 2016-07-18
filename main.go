package main

import (
	"github.com/wangbai/peeper/httpserv"
)

func main() {
	//dis := &monitor.Discovery {}
	//d := httpserv.NewDiscoveryResource(dis)
	httpserv.Start(11111)
}
