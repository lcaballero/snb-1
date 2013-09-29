package caching

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func Test_CacheEntries_Setup(t *testing.T) {

	fmt.Println("os.Args[0]", os.Args[0])

	for i, e := range os.Args {
		fmt.Println(i, e)
	}

	var config string
	flag.StringVar(&config, "config-file", "default", "Configuration file location.")
	flag.Parse()

	for i, e := range flag.Args() {
		fmt.Println(i, e)
	}

	fmt.Println("config:", config)

	fmt.Println("CWD: ", os.Args[0])
	hasSql(t, CacheEntries.AddUserToGroup)
}

func hasSql(t *testing.T, c *CacheEntry) {
	if c.Path == "" {
		t.Error("Cache Entry provided doesn't have a path.")
	}

	if c.Script == "" {
		t.Error("CacheEntry doesn't contain any sql code.")
	}

	if c.Err != nil {
		t.Error("An error occured while reading a CacheEntry with path: '" + c.Path + "'")
	}
}
