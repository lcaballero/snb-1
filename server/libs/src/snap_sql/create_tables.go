package snap_sql

import (
	"fmt"
	"io/ioutil"
	"sql_utils"
)

// ---------------------- Create Board Table ---------------------- //

func CreateBoardTable() {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createBoardTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		_, err := sql_utils.GetConnection().Exec(string(sql))

		if err != nil {
			fmt.Println(err)
		}
	}
}

// ---------------------- Create Game Table ---------------------- //

func CreateGameTable() {

	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createGameTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		_, err := sql_utils.GetConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(result)
	}
}

// ---------------------- Create User Table ---------------------- //
func CreateUserTable() {

	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createUserTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		_, err := sql_utils.GetConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(result)
	}
}

// ---------------------- Create Group Table ---------------------- //

func CreateGroupsTable() {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createGroupsTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		_, err := sql_utils.GetConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(result)
	}
}

// ---------------------- Create User to Group table ---------------------- //

func CreateUserToGroupTable() {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createUserToGroupTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		_, err := sql_utils.GetConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(result)
	}
}
