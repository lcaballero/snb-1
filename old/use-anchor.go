package main

import (
	"anchored"
	"fmt"
)

type User struct {
	anchored.Anchor
	anchor anchored.Map
}

func (u *User) GetAnchor() anchored.Map {
	return u.anchor
}

func (u *User) SetAnchor(m anchored.Map) {
	u.anchor = m
}

func main() {
	fmt.Println("hello")
}
