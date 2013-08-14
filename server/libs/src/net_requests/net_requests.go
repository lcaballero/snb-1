package net_requests

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func MakeRequest(path string) (string, error) {

	res, err := http.Get(path)

	if err != nil {
		fmt.Println(err)
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	return string(robots), err
}
