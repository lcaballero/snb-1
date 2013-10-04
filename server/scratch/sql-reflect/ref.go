package main

import (
	"fmt"
	"reflect"
)

type DataPoints map[string]interface{}
type ItemProvider func() interface{}
type Saver func(Any)
type Any interface{}

type Base struct {
	Id        string
	Status    int
	UpdatedOn string
	UpdatedBy string
}

type Struct struct {
	Base
	Name, Email string
	Age         int
}

func main() {
	compareStructs()
}

func compareStructs() {
	a := &Struct{Name: "Lucas"}
	b := &Struct{Name: "Lucas"}

	if *a == *b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a != b")
	}
}

func Set(target Any, field string, val Any) {

	ptrValue := reflect.ValueOf(target)

	if !ptrValue.IsValid() {
		return
	}

	value := reflect.Indirect(ptrValue)

	if !value.IsValid() {
		return
	}

	f := value.FieldByName(field)

	if f.IsValid() {
		f.Set(reflect.ValueOf(val))
	}
}

func FromMap(provider ItemProvider, data DataPoints) Any {
	item := provider()

	for k, v := range data {
		Set(item, k, v)
	}

	return item
}

func FromMaps(provide ItemProvider, saver Saver, data []DataPoints) {

	for _, m := range data {
		val := FromMap(provide, m)
		saver(val)
	}
}
