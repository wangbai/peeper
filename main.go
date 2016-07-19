package main

import (
	"github.com/wangbai/peeper/config"
	"github.com/wangbai/peeper/httpserv"
	_ "github.com/wangbai/peeper/monitor"
)

func main() {
	config.Build("config/")

	server := httpserv.NewServer()
	server.Start()
}
