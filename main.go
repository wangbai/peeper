package main

import (
	"flag"
	"os"

	"github.com/wangbai/peeper/config"
	"github.com/wangbai/peeper/httpserv"
	_ "github.com/wangbai/peeper/monitor"
)

var configDir string

func parseCmdLine() {
	flag.StringVar(&configDir, "d", "", "directory for config files")
	flag.Parse()

	if configDir == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	parseCmdLine()

	config.Build(configDir)

	server := httpserv.NewServer()
	server.Start()
}
