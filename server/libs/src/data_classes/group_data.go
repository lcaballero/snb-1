package data_classes

import(
	"fmt"
	"time"
	//enc "json_helpers"
)

/* ---------------------- Group Data ---------------------- */

type GroupData struct {
	Anchor
}

func (group GroupData) Id() string {
	fmt.Println("--------", group.GetProp("id").(string))
	return group.GetProp("id").(string)
}

func (group GroupData) GroupName() string {
	return group.GetProp("group_name").(string)
}

func (group GroupData) Description() string {
	return group.GetProp("group_desc").(string)
}

func (group GroupData) DateAdded() time.Time {
	return group.GetProp("date_added").(time.Time)
}