package snap_sql

import (
	"fmt"
	"io/ioutil"
	"sql_utils"
)

// ---------------------- Create Tile Table ---------------------- //

func CreateTileTable() {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createTileTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		_, err := sql_utils.GetConnection().Exec(string(sql))

		if err != nil {
			fmt.Println(err)
		}
	}
}

// ---------------------- Create Criteria Table ---------------------- //

func CreateCriteriaTable() {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createCriteriaTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		_, err := sql_utils.GetConnection().Exec(string(sql))

		if err != nil {
			fmt.Println(err)
		}
	}
}

// ---------------------- Create Board Table ---------------------- //

func CreateBoardTable() (err error) {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createBoardTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		_, err := sql_utils.GetConnection().Exec(string(sql))

		if err != nil {
			fmt.Println(err)
		}
	}

	return err
}

// ---------------------- Create Game Table ---------------------- //

func CreateGameTable() (err error) {

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

	return err
}

// ---------------------- Create User Table ---------------------- //
func CreateUserTable() (err error) {

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

	return err
}

// ---------------------- Create Group Table ---------------------- //

func CreateGroupsTable() (err error) {
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

	return err
}

// ---------------------- Create User to Group table ---------------------- //

func CreateUserToGroupTable() (err error) {
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

	return err
}
