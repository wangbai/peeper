package main

import (
	"flag"
	"os"
	"runtime"

	_ "github.com/wangbai/peeper/monitor"
	"github.com/wangbai/peeper/server"
	_ "github.com/wangbai/peeper/snowflake"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var configDir string
var port int
var dryrun bool

func parseCmdLine() {
	flag.StringVar(&configDir, "config_dir", "", "directory for config files")
	flag.IntVar(&port, "port", 16722, "local server port, default: 16722")
	flag.BoolVar(&dryrun, "dryrun", false, "dryrun for checking config, default: false")
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
