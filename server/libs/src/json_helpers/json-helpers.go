package json_helpers

import (
	"encoding/json"
)

func ToJson(v interface{}) (s string) {
	j, _ := json.Marshal(v)
	s = string(j)
	return s
}

func ToIndentedJson(v interface{}, prefix, indent string) (s string) {
	j, _ := json.MarshalIndent(v, prefix, indent)
	s = string(j)
	return s
}

func FromJson(jsonBlob string, val interface{}) error {
	bytes := []byte(jsonBlob)
	err := json.Unmarshal(bytes, val)
	return err
}