package main

import (
	"fmt"
	//"builtin"
	_ "github.com/bmizerany/pq"
	"io/ioutil"
	"requests"
	"snap_sql"
	"sql_utils"
	"uuid"
	//"reflect"
	//"sql_text"
)

// ---------------------- User Functions ---------------------- //

func tableExists(dbName, tableName string) bool {

	sql, err := ioutil.ReadFile(sql_utils.FilePath + "tableExists.sql")

	if err != nil {
		fmt.Println(err)
		return false
	}

	rows, err := sql_utils.GetConnection().Query(string(sql), dbName, tableName)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		m := requests.ToMapping(rows)
		hasLen := len(m) > 0

		if hasLen {
			fmt.Println("Table Exists: ", tableName)
		} else {
			fmt.Println("Table Does NOT Exists: ", tableName)
		}

		return hasLen
	}
}

/* ------------------------- Has Table methods ------------------------- */

func hasUserTable() bool {
	has_table := tableExists("snb", "_user")
	return has_table
}

func hasGroupTable() bool {
	has_table := tableExists("snb", "socialgroup")
	return has_table
}

func hasUserToGroupTable() bool {
	has_table := tableExists("snb", "usertogroup")
	return has_table
}

func hasGameTable() bool {
	return tableExists("snb", "game")
}

/* ------------------------- Main ------------------------- */

func main() {

	if !hasUserTable() {
		fmt.Println("Creating User Table")
		snap_sql.CreateUserTable()
	}

	if !hasGroupTable() {
		fmt.Println("Create Group table")
		snap_sql.CreateGroupsTable()
	}

	if !hasUserToGroupTable() {
		fmt.Println("Create UserToGroup table")
		snap_sql.CreateUserToGroupTable()
	}

	if !hasGameTable() {
		fmt.Println("Create Game table")
		snap_sql.CreateGameTable()
	}

	/* ------------------------- Create Group ------------------------- */

	hasGlobalGroup, _ := snap_sql.HasGroup("global_group")
	if !hasGlobalGroup {
		globalGroupUuid := uuid.New()
		group_status, _ := snap_sql.CreateGroup("global_group", "group that contains every user", globalGroupUuid)
		fmt.Println("Create Group: ", group_status.Msg)
	}

	/* ------------------------- Create User ------------------------- */

	usr := "d333"
	pw := "ro"

	status, _ := snap_sql.CreateUser(usr, pw)

	fmt.Println("Create User: ", status.Msg)

	breweryGroup := "Breweries"
	hasBreweryGroup, _ := snap_sql.HasGroup(breweryGroup)
	if !hasBreweryGroup {
		myUser, _ := snap_sql.ReadUserByEmail(usr)

		group_status, _ := snap_sql.CreateGroup(breweryGroup, "Breweries in Boulder", myUser[0].Id)
		fmt.Println("Create Group: ", breweryGroup, group_status.Msg)
	}

	/* ------------------------- Read User By Email ------------------------- */
	//userByEmail, err := readUserByEmail("Ryan")

	fmt.Println()
	fmt.Println()

	// for i := 0; i < len(userByEmail); i++ {
	// 	fmt.Println(enc.ToIndentedJson(userByEmail[i].PrintAll(), "", "  "));
	// }

	/* ------------------------- Read All Users ------------------------- */
	// allUsers, err := snap_sql.ReadAllUsers()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println()
	// fmt.Println()

	// for i := 0; i < len(allUsers); i++ {
	// 	fmt.Println(enc.ToIndentedJson(allUsers[i].PrintAll(), "", "  "));
	// }

}
