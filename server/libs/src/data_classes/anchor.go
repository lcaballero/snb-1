package data_classes

import (
	"fmt"
	//enc "json_helpers"
)

/* ---------------------- Anchor Dictionary..ish ---------------------- */

type Anchor struct {
	refMap map[string]interface{}
}

func (anchor *Anchor) SetMap(m map[string]interface{}){
	anchor.refMap = m;
}

func (anchor *Anchor) GetProp(reqField string) interface{} {
	oField, ok := anchor.refMap[reqField]

	if !ok {
		fmt.Println("Anchor refMap err: ", ok, reqField)
		return ""
	}

	return oField
}
