package main

import (
	"flag"
	"os"

	"github.com/golang/glog"

	"github.com/ReflecBeatCustom/haereticus/cmd/server"
	ver "github.com/ReflecBeatCustom/haereticus/pkg/version"
)

var (
	configFile string
	version    bool
	// GitTag ...
	GitTag string = "v1.0.0"
	// BuildTime ...
	BuildTime string = "2020-08-02T--:00:00+0800"
)

func main() {
	flag.StringVar(&configFile, "conf", "../etc/server.toml", "toml config file to read from")
	flag.BoolVar(&version, "version", false, "show version")
	flag.Parse()
	defer glog.Flush()

	if version || (len(os.Args) > 1 && os.Args[1] == "version") {
		ver.PrintVersionAndExit()
	}

	server := server.NewServer(configFile)
	server.Start()
}
