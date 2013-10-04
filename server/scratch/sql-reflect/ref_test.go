package main

import (
	"testing"
)

func Test_Set(t *testing.T) {

	s := &Struct{Name: "Ryan"}

	Set(s, "Name", "Lucas")

	if s.Name != "Lucas" {
		t.Error("Set() didn't update the Name.")
	}
}

func Test_Set_Mismatched_Name(t *testing.T) {

	s := &Struct{Name: "Ryan"}

	Set(s, "Bad_Name", "Lucas")
}

func Test_Comparing_Struct(t *testing.T) {
	a := Struct{}
	a.Name = "Lucas"
	a.Id = "id-1"

	b := Struct{}
	b.Name = "Lucas"
	b.Id = "id-1"

	if a != b {
		t.Error("Comparing structs with anonymous Base fields.")
	}
}

func Test_Null_To_Set(t *testing.T) {
	Set(nil, "Bad_Name", "Bad_Value")
}

func provide() interface{} {
	return &Struct{}
}

func Test_From_Map(t *testing.T) {

	a := &Struct{
		Name: "Lucas",
		Base: Base{
			Id: "id-2",
		},
	}

	m := DataPoints{
		"Name": a.Name,
		"Id":   a.Id,
	}

	rv := FromMap(provide, m)

	item := rv.(*Struct)

	hasMapped := item.Name == a.Name && item.Id == a.Id

	if !hasMapped {
		t.Error("Did not map name from mapping")
	}
}

func Test_From_Maps(t *testing.T) {

	a := &Struct{
		Name: "Ryan",
		Base: Base{Id: "id-0"},
	}

	b := &Struct{
		Name: "Lucas",
		Base: Base{Id: "id-1"},
	}

	m := []DataPoints{
		DataPoints{"Name": a.Name, "Id": a.Id},
		DataPoints{"Name": b.Name, "Id": b.Id},
	}

	vals := make([]*Struct, 0)

	saver := func(a Any) {
		vals = append(vals, a.(*Struct))
	}

	FromMaps(provide, saver, m)

	x := find(a.Id, vals)
	y := find(b.Id, vals)

	if len(vals) != len(m) {
		t.Error("The resulting data count doesn't match the data provided.")
	}

	matches := *a != *x && *b != *y

	if matches {
		t.Error("Structs out don't match input Structs")
	}
}

func find(id string, items []*Struct) *Struct {
	for _, e := range items {
		if e.Id == id {
			return e
		}
	}
	return nil
}
