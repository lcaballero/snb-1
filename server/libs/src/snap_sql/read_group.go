package snap_sql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"data_classes"
	"sql_utils"
)

// ---------------------- Read Group Functions ---------------------- //

func HasGroup(groupName string) (bool, error) {
	groups, err := ReadGroup(groupName)

	if err != nil {
		fmt.Println(err)
		return true, err // TODO: should this be true or false?
	} else if len(groups) > 0 {
		return true, err
	} else {
		return false, err
	}
}

func ReadGroup(group_name string) ([]data_classes.GroupData, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "readGroup.sql");

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		rows, err := sql_utils.GetConnection().Query(string(sql), group_name)

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			return processGroup(rows, err)
		}
	}
}

func processGroup(sqlRows *sql.Rows, err error) ([]data_classes.GroupData, error) {
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		mappedRows := sql_utils.ToSqlMap(sqlRows)

		groups := make([]data_classes.GroupData, len(mappedRows))

		for i, v := range mappedRows {
			// anchor := data_classes.Anchor{}
			// anchor.SetMap(v)
			groups[i] = data_classes.GroupData{Anchor:data_classes.Anchor{}}
			groups[i].SetMap(v)
		}

		//fmt.Println("group[0]:", groups[0].GetProp("group_name"))
		return groups, nil
	}
}
