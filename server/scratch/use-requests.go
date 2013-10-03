package main

import (
	"fmt"
	//"builtin"
	"data_classes"
	_ "github.com/bmizerany/pq"
	enc "json_helpers"
	"snap_sql"
	_ "sql_utils"
	"sql_utils/caching"
	"sql_utils/codes"
	"strconv"
	"uuid"
	//"reflect"
	//"sql_text"
)

func main() {

	caching.LoadSqlScripts()
	snap_sql.SetupTables()

	/* ------------------------- Create Group ------------------------- */

	hasGlobalGroup, _ := snap_sql.HasGroup("global_group")

	if !hasGlobalGroup {
		globalGroupUuid := uuid.New()
		group_status, _ := snap_sql.CreateGroup(globalGroupUuid, "global_group", "group that contains every user", globalGroupUuid)
		fmt.Println("Create Group: ", group_status)
	} else {
		fmt.Println("Has Group: ", "global_group")
	}

	/* ------------------------- Create User ------------------------- */

	usr := "d333"
	pw := "ro"

	status, _ := snap_sql.CreateUser(usr, pw)

	fmt.Println("Create User: ", status)

	myUser, _ := snap_sql.ReadUserByEmail(usr)

	/* ------------------------- Create a Group ------------------------- */

	breweryGroupName := "Breweries"
	hasBreweryGroup, _ := snap_sql.HasGroup(breweryGroupName)
	if !hasBreweryGroup {
		groupUuid := uuid.New()

		group_status, _ := snap_sql.CreateGroup(groupUuid, breweryGroupName, "Breweries in Boulder", myUser[0].Id)
		fmt.Println("Create Group: ", breweryGroupName, group_status)
	} else {
		fmt.Println("Has Group: ", breweryGroupName)
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

		/* ------------------------- Create a board ------------------------- */

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

			fmt.Println("Create Board: ", boardName, board_status)
		}
	}

	/* ------------------------- Create Criteria ------------------------- */

	for i := 0; i < 25; i++ {
		criteriaUuid := uuid.New()

		crit_status, err := snap_sql.CreateCriteria(
			criteriaUuid,
			"crit_"+strconv.Itoa(i),
		)

		if err != nil {
			fmt.Println(err)
		}
		if crit_status == codes.Success {
		}
		//crit_status.String()
	}

	userBoards, err := snap_sql.ReadUsersBoards(myUser[0].Id)

	if err != nil {
		fmt.Println("Read User Boards err: ", err)
	} else {
		fmt.Println()
		fmt.Println("Read User Boards... ", myUser[0].Id)
		fmt.Println(enc.ToIndentedJson(userBoards, "", "  "))
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
