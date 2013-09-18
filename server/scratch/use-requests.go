package main

import (
	"fmt"
	//"builtin"
	"data_classes"
	_ "github.com/bmizerany/pq"
	"io/ioutil"
	enc "json_helpers"
	"requests"
	"snap_sql"
	"sql_utils"
	"uuid"
	//"reflect"
	//"sql_text"
)

// ---------------------- User Functions ---------------------- //

func dropTable(tableName string) bool {

	sql := "drop table if exists " + tableName

	fmt.Println(sql)

	result, err := sql_utils.GetConnection().Exec(string(sql))

	if err != nil {
		fmt.Println("querying err: ", err)
		return false
	}

	fmt.Println("result: ", result)

	return true
}

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
		group_status, _ := snap_sql.CreateGroup(globalGroupUuid, "global_group", "group that contains every user", globalGroupUuid)
		fmt.Println("Create Group: ", group_status.Msg)
	}

	/* ------------------------- Create User ------------------------- */

	usr := "d333"
	pw := "ro"

	status, _ := snap_sql.CreateUser(usr, pw)

	fmt.Println("Create User: ", status.Msg)

	/* ------------------------- Create a Group ------------------------- */

	breweryGroupName := "Breweries"
	hasBreweryGroup, _ := snap_sql.HasGroup(breweryGroupName)
	if !hasBreweryGroup {
		myUser, _ := snap_sql.ReadUserByEmail(usr)
		groupUuid := uuid.New()

		group_status, _ := snap_sql.CreateGroup(groupUuid, breweryGroupName, "Breweries in Boulder", myUser[0].Id)
		fmt.Println("Create Group: ", breweryGroupName, group_status.Msg)
	}

	/* ------------------------- Create a game ------------------------- */

	// read the group and find its Id
	hasBreweryGroup, _ = snap_sql.HasGroup(breweryGroupName)
	var breweryGroup []data_classes.GroupData

	if hasBreweryGroup {
		breweryGroup, _ = snap_sql.ReadGroup(breweryGroupName)

		// Create a game and assign it to breweryGroup
		gameUuid := uuid.New()
		createGame_status, _ := snap_sql.CreateGame(
			gameUuid, breweryGroup[0].Id, "Boulder Breweries", "Have a brew with a brewer")

		fmt.Println("Create Game: ", "Boulder Breweries", createGame_status)

		fmt.Println("Read Game from Id: ...")
		readGameFromId, _ := snap_sql.ReadGameFromId(gameUuid)
		fmt.Println(enc.ToIndentedJson(readGameFromId, "", "  "))

		fmt.Println("Read Game from Name: ...")
		readGameFromName, _ := snap_sql.ReadGameFromName("Boulder Breweries")
		fmt.Println(enc.ToIndentedJson(readGameFromName, "", "  "))
	}

	// read all games in a group

	fmt.Println("Read all games in group: ...")
	allGameInGroup, _ := snap_sql.ReadAllGames(breweryGroup[0].Id)

	for i := 0; i < len(allGameInGroup); i++ {
		fmt.Println(enc.ToIndentedJson(allGameInGroup[i], "", "  "))
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
