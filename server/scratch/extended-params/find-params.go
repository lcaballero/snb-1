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

