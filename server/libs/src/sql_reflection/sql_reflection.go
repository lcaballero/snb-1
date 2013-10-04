package sql_reflection

import (
	"reflect"
	"strings"
)

func Set(target interface{}, field string, val interface{}) {

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

func FromMap(item interface{}, data map[string]interface{}) {
	for k, v := range data {
		field := ColunnToFieldName(k)
		Set(item, field, v)
	}
}

func FromMaps(ptrs []interface{}, data []map[string]interface{}) {
	for i, m := range data {
		FromMap(ptrs[i], m)
	}
}

func Fields(a interface{}) []string {

	s := make([]string, 0)

	t := reflect.TypeOf(a)
	count := t.NumField()

	for i := 0; i < count; i++ {
		f := t.Field(i)
		s = append(s, f.Name)
	}

	return s
}

func Capitalize(name string) string {
	if len(name) == 0 {
		return name
	}

	s := strings.ToUpper(name[0:1]) + name[1:]

	return s
}

func ColunnToFieldName(col string) string {
	parts := strings.Split(col, "_")
	n := len(parts)
	buf := make([]string, 0)

	for i := 0; i < n; i++ {
		p := parts[i]
		c := Capitalize(p)
		buf = append(buf, c)
	}

	name := strings.Join(buf, "")

	return name
}
