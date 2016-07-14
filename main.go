package main

import (
	"fmt"
	"time"

	"github.com/wangbai/peeper/service"
)

func main() {
	serv := service.NewService("wangbai",
		[]service.Node{
            {Address: "www.baidu.com:80", Interval: 2 * time.Second},
            {Address: "www.sohu.com:80", Interval: 2 * time.Second},
            {Address: "www.wangbai.com:80", Interval: 2 * time.Second},
            {Address: "gitlab.greencheng.com:9090", Interval: 2 * time.Second},
        },
	)

	serv.StartMonitor()

	for {
		time.Sleep(2 * time.Second)

		fmt.Println(serv)
	}
}
