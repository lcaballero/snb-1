package main

import (
	"os"
	"path"
	"path/filepath"
)

func FindConfigFile(file string) (string, bool) {

	exists := func(file string) bool {
		_, err := os.Stat(file)
		exists := file != "" && !os.IsNotExist(err)
		return exists
	}

	abs, _ := filepath.Abs(file)
	dir := path.Dir(abs)
	found := exists(abs)

	for !found {

		parent := dir + "/../"
		dir = path.Clean(parent)
		abs = path.Join(dir, file)

		found = exists(abs)

		if dir == "/" {
			break
		}
	}

	return abs, found
}
