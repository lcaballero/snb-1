package main

import (
	"flag"
	"fmt"
)

type commandFlags struct {
	config_file     string
	help            bool
	show_parameters bool
}

func main() {

	cf := &commandFlags{}

	flag.StringVar(&cf.config_file, "config-file", "", "Absolute path to the root of the application where the config.js should reside.")
	flag.BoolVar(&cf.help, "help", false, "Show this help message.")
	flag.BoolVar(&cf.show_parameters, "show-parameters", false, "Show the parameters sent to the command, but don't run the command.")

	flag.Parse()

	if cf.help {
		flag.Usage()
		return
	}

	if cf.show_parameters {
		fmt.Println("config_file: ", cf.config_file)
		fmt.Println("help: ", cf.help)
		fmt.Println("show_parameters: ", cf.show_parameters)
		return
	}
}
