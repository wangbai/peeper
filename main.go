package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/wangbai/peeper/config"
	"github.com/wangbai/peeper/httpserv"
	_ "github.com/wangbai/peeper/monitor"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var configDir string
var dryrun bool

func parseCmdLine() {
	flag.StringVar(&configDir, "d", "", "directory for config files")
	flag.BoolVar(&dryrun, "dryrun", false, "dryrun for checking config")
	flag.Parse()

	if configDir == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	parseCmdLine()

	config.Build(configDir)

    if dryrun {
	    server := httpserv.NewServer()
	    server.Start()
    }
}
