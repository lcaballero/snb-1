package snap_sql

import(
	"fmt"
	"io/ioutil"
	"sql_utils"
)

func CreateUserTable() {

	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createUserTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		result, err := sql_utils.GetConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}

// ---------------------- Group Functions ---------------------- //

func CreateGroupsTable() {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createGroupsTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		result, err := sql_utils.GetConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}

func CreateUserToGroupTable() {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createUserToGroupTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		result, err := sql_utils.GetConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}