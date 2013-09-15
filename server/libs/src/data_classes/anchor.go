package data_classes

import(
	"fmt"
	enc "json_helpers"
)

/* ---------------------- Anchor Dictionary..ish ---------------------- */

type Anchor struct {
	refMap map[string]interface{}
}

func (anchor Anchor) SetMap(m map[string]interface{}){
	anchor.refMap = m;

	fmt.Println("---=================---")
	fmt.Println(enc.ToIndentedJson(anchor.refMap, "", "  "))
}

func (anchor Anchor) GetProp(reqField string) interface{} {
	oField, ok := anchor.refMap[reqField]

	fmt.Println("------")
	fmt.Println(enc.ToIndentedJson(anchor.refMap, "", "  "))
	if !ok {
		fmt.Println("Anchor refMap err: ", ok, reqField)
		return ""
	}

	return oField
	/*
	// TODO: need to abstract the type cast so we can use
	// other type such as int.
	field, ok := oField.(string)
	
	if ok {
		return field
	} else {
		fmt.Println("Error: Unable to convert %v to a string", reqField)
		return ""
	}
	return field, ok
	*/
}