package session_manager

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var Conn *redis.Conn = nil

func GetConn() redis.Conn {

	if Conn == nil {
		conn, err := redis.DialTimeout(
			"tcp", "127.0.0.1:6379", 0, 1*time.Second, 1*time.Second)
		if err != nil {
			fmt.Println(err)
		}
		Conn = &conn
	}

	return *Conn
}
