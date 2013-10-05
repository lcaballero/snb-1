package rt_config

import "os"

func exists(file string) bool {
	_, err := os.Stat(file)
	exists := file != "" && !os.IsNotExist(err)
	return exists
}
