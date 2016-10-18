package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/wangbai/peeper/server"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var configDir string
var port int
var dryrun bool

func parseCmdLine() {
	flag.StringVar(&configDir, "config_dir", "", "directory for config files")
	flag.IntVar(&port, "port", 0, "local server port")
	flag.BoolVar(&dryrun, "dryrun", false, "dryrun for checking config")
	flag.Parse()

	if configDir == "" || port == 0 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	parseCmdLine()

	server.LoadModule(configDir)

	if !dryrun {
	    server.NewServer(port).Start()
	}
}
