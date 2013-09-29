package main

import (
	"flag"
	"fmt"
)

func main() {
	var config string
	flag.StringVar(&config, "config-file", "Default", "Config file location.")
	flag.Parse()

	fmt.Println("Config:", config)
}
