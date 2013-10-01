package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	flag := "--config-file="
	val := findFlag(flag)

	fmt.Println(flag + val)
}

func FindFlag(flag string) string {

	val := ""

	for _, e := range os.Args {

		hasPrefix := strings.HasPrefix(e, flag)

		if hasPrefix {
			val = e[len(flag):]
		}
	}

	return val
}
