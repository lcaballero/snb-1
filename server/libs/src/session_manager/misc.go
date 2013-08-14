package session_manager

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"session_manager/hashes"
)

func ReadHashPairs(key string) (vals map[string]string, res interface{}, err error) {

	res, err = GetConn().Do(hashes.HGETALL, key)

	if err != nil {
		fmt.Println(err)
		return
	}

	return ReadPairs(res)
}

func ReadPairs(reply interface{}) (vals map[string]string, res interface{}, err error) {

	rows, ok := reply.([]interface{})

	if !ok {
		fmt.Println("Result not an []interface{}.")
		return vals, reply, err
	}

	ln := len(rows)

	if ln%2 != 0 {
		fmt.Println("len(rows) is not even: ", ln)
		return vals, reply, err
	}

	count := ln / 2
	vals = make(map[string]string)

	err = nil

	for i := 0; i < count; i++ {
		n := i * 2
		k, err := redis.Bytes(rows[n], err)
		v, err := redis.Bytes(rows[n+1], err)

		if err != nil {
			fmt.Println("Couldn't convert via redis.Bytes(row)")
		}

		vals[string(k)] = string(v)
	}

	return vals, reply, err
}
