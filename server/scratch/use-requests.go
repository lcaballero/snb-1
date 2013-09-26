package main

import (
	"fmt"
	//"builtin"
	"data_classes"
	_ "github.com/bmizerany/pq"
	enc "json_helpers"
	"snap_sql"
	_ "sql_utils"
	"uuid"
	//"reflect"
	//"sql_text"
)

func main() {

	if !snap_sql.HasUserTable() {
		fmt.Println("Creating User Table...")
		snap_sql.CreateUserTable()
	}

	if !snap_sql.HasGroupTable() {
		fmt.Println("Create Group table...")
		snap_sql.CreateGroupsTable()
	}

	if !snap_sql.HasUserToGroupTable() {
		fmt.Println("Create UserToGroup table...")
		snap_sql.CreateUserToGroupTable()
	}

	if !snap_sql.HasGameTable() {
		fmt.Println("Create Game table...")
		snap_sql.CreateGameTable()
	}

	if !snap_sql.HasBoardTable() {
		fmt.Println("Create Board table...")
		snap_sql.CreateBoardTable()
	}

	if !snap_sql.HasCriteriaTable() {
		fmt.Println("Create Criteria table...")
		snap_sql.CreateCriteriaTable()
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

	myUser, _ := snap_sql.ReadUserByEmail(usr)

	/* ------------------------- Create a Group ------------------------- */

	breweryGroupName := "Breweries"
	hasBreweryGroup, _ := snap_sql.HasGroup(breweryGroupName)
	if !hasBreweryGroup {
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

		createNewGame := false

		if createNewGame {
			// Create a game and assign it to breweryGroup
			gameUuid := uuid.New()
			createGame_status, _ := snap_sql.CreateGame(
				gameUuid, breweryGroup[0].Id, "Boulder Breweries", "Have a brew with a brewer")

			fmt.Println("Create Game: ", "Boulder Breweries", createGame_status)

			fmt.Println("Read Game from Id: ...")
			readGameFromId, _ := snap_sql.ReadGameFromId(gameUuid)
			fmt.Println(enc.ToIndentedJson(readGameFromId, "", "  "))
		}

		fmt.Println()
		fmt.Println("Read Game in group from Name: ...")
		readGameInGroupFromName, _ := snap_sql.ReadGameInGroupFromName(
			breweryGroup[0].Id, "Boulder Breweries")
		fmt.Println(enc.ToIndentedJson(readGameInGroupFromName, "", "  "))

		createNewBoard := false

		if createNewBoard {
			boardUuid := uuid.New()
			boardName := "I drink too much"

			board_status, _ := snap_sql.CreateBoard(
				boardUuid,
				readGameInGroupFromName[0].Id,
				myUser[0].Id,
				boardName,
				1)

			// for i := 0; i < 25; i++ {
			// 	criteriaUuid := uuid.New()

			// 	_, _ := snap_sql.CreateCriteria(
			// 		criteriaUuid,
			// 		"crit_" + i
			// 	)
			// }
			fmt.Println("Create Board: ", boardName, board_status)
		}
	}

	/*
		fmt.Println()
		fmt.Println("Read Game from Name: ...")
		readGameFromName, _ := snap_sql.ReadGameFromName("Boulder Breweries")
		fmt.Println(enc.ToIndentedJson(readGameFromName, "", "  "))
	*/
	// read all games in a group
	/*
		fmt.Println("Read all games in group: ...")
		allGameInGroup, _ := snap_sql.ReadAllGames(breweryGroup[0].Id)

		for i := 0; i < len(allGameInGroup); i++ {
			fmt.Println(enc.ToIndentedJson(allGameInGroup[i], "", "  "))
		}
	*/

	/* ------------------------- Read User By Email ------------------------- */

	//userByEmail, err := readUserByEmail("Ryan")
	/*
		fmt.Println()

		for i := 0; i < len(userByEmail); i++ {
			fmt.Println(enc.ToIndentedJson(userByEmail[i].PrintAll(), "", "  "));
		}
	*/

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
