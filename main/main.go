package main

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/infra/conf/serial"
	_ "github.com/xtls/xray-core/main/distro/all"
)

func main() {
	fmt.Println("hello world")
	//configFiles := cmdarg.Arg{path.Join("/Users/puma/Desktop/code/xray-core-lib/main", "config.json")}
	//c, err := core.LoadConfig("auto", configFiles)
	filename := "/Users/puma/Desktop/code/xray-core-lib/main/config.json"
	//var data []byte
	data , _ := os.ReadFile(filename)
	out := bytes.NewBuffer(data)
	c, _ := serial.LoadJSONConfig(out)
	server, err := core.New(c)
	//_ = server.GetFeature(v2stats.ManagerType()).(v2stats.Manager)
	server.Start()
	if err != nil {
		println("error!")
	}

	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		<-osSignals
	}
}
