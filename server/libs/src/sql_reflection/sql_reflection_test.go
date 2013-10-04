package sql_reflection

import (
	"sql_utils/caching"
	"testing"
)

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

func init() {
	caching.LoadSqlScripts()
}

func Test_Compare_Structs(t *testing.T) {
	a := &Struct{Name: "Lucas"}
	b := &Struct{Name: "Lucas"}

	if *a != *b {
		t.Error("a != b")
	}
}

func Test_Captialize_Normal(t *testing.T) {

	m := map[string]string{
		"date": "Date",
		"":     "",
		"d":    "D",
		"123d": "123d",
		"  ":   "  ",
	}

	for k, v := range m {
		c := Capitalize(k)
		if c != v {
			t.Error(
				"Not capitalized: ", k,
				" result: ", c,
				" should be: ", v)
		}
	}
}

func Test_Column_To_Field(t *testing.T) {
	m := map[string]string{
		"date_added":                    "DateAdded",
		"_user":                         "User",
		"updated_on":                    "UpdatedOn",
		"updated_by":                    "UpdatedBy",
		"_field_with_many_underscores_": "FieldWithManyUnderscores",
	}

	for k, v := range m {
		field := ColunnToFieldName(k)
		if field != v {
			t.Error("Field doesn't match: ", k, ":", field)
		}
	}
}

func Test_Fields(t *testing.T) {
	s := Struct{}
	fields := Fields(s)

	names := []string{
		"Id", "Name", "Email", "Age",
		"Status", "UpdatedOn", "UpdatedBy",
	}

	if len(fields) == len(names) {
		t.Error("The number of fields aren't equal.")
	}
}

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

func NewStruct() *Struct {
	return &Struct{}
}

func Test_From_Map(t *testing.T) {

	a := &Struct{
		Name: "Lucas",
		Base: Base{
			Id: "id-2",
		},
	}

	m := map[string]interface{}{
		"Name": a.Name,
		"Id":   a.Id,
	}

	item := NewStruct()
	FromMap(item, m)

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

	m := []map[string]interface{}{
		map[string]interface{}{"Name": a.Name, "Id": a.Id},
		map[string]interface{}{"Name": b.Name, "Id": b.Id},
	}

	count := len(m)
	structs := make([]*Struct, count)
	ptrs := make([]interface{}, count)

	for i := 0; i < count; i++ {
		structs[i] = &Struct{}
		ptrs[i] = structs[i]
	}

	FromMaps(ptrs, m)

	x := find(a.Id, structs)
	y := find(b.Id, structs)

	if len(structs) != len(m) {
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
