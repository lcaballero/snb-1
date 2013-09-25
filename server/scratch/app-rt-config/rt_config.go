package main

import (
	"fmt"
	"rt_config"
)

func main() {

	cf := rt_config.LoadFromCommandLine()

	fmt.Println(cf)

	rt_config.DumpConfigFile()
}
