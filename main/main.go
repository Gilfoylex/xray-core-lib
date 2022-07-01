package main

import (
	"fmt"

	"github.com/xtls/xray-core/core"
)

func main() {
	fmt.Println("hello world")
	configFiles := [...]string{"D:/xray-config/config.json"}
	c, err := core.LoadConfig("json", configFiles)
	core.New(c)
	if err != nil {
		println("error!")
	}
}
