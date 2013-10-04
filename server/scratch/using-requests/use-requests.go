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

func init() {
	caching.LoadSqlScripts()
}

const (
	usr                = "d333"
	pw                 = "ro"
	global_group_name  = "global_group"
	global_group_desc  = "group that contains every user"
	brewery_group_name = "Breweries"
)

func createUser() []*data_classes.UserProfile {

	status, _ := snap_sql.CreateUser(usr, pw)

	fmt.Println("Create User: ", status)

	myUser, _ := snap_sql.ReadUserByEmail(usr)

	return myUser
}

func createGlobalGroup() {

	hasGlobalGroup, _ := snap_sql.HasGroup("global_group")

	if !hasGlobalGroup {

		globalGroupUuid := uuid.New()
		group_status, _ := snap_sql.CreateGroup(
			globalGroupUuid,
			global_group_name,
			global_group_desc,
			globalGroupUuid)

		fmt.Println("Create Group: ", group_status)

	} else {

		fmt.Println("Has Group: ", "global_group")
	}
}

func createBreweriesGroup(userId string) {

	hasBreweryGroup, _ := snap_sql.HasGroup(brewery_group_name)
	if !hasBreweryGroup {
		groupUuid := uuid.New()

		group_status, _ := snap_sql.CreateGroup(groupUuid, brewery_group_name, "Breweries in Boulder", userId)
		fmt.Println("Create Group: ", brewery_group_name, group_status)
	} else {
		fmt.Println("Has Group: ", brewery_group_name)
	}
}

func createGame(user *data_classes.UserProfile) {
	/* ------------------------- Create a game ------------------------- */

	// read the group and find its Id
	hasBreweryGroup, _ := snap_sql.HasGroup(brewery_group_name)
	var breweryGroup []data_classes.GroupData

	if hasBreweryGroup {
		breweryGroup, _ = snap_sql.ReadGroup(brewery_group_name)

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
		//fmt.Println(enc.ToIndentedJson(readGameInGroupFromName, "", "  "))

		/* ------------------------- Create Criteria ------------------------- */

		for i := 0; i < 25; i++ {
			criteriaUuid := uuid.New()

			crit_status, err := snap_sql.CreateCriteria(
				criteriaUuid,
				"crit_"+strconv.Itoa(i),
			)

			/* ------------------------- Add Criteria to Game ------------------------- */

			if err != nil {
				fmt.Println(err)
			}
			if crit_status == codes.Success {
			}

			critToGameUuid := uuid.New()
			crit_to_game, err := snap_sql.AddCriteriaToGame(
				critToGameUuid, readGameInGroupFromName[0].Id, criteriaUuid, 1, 1)

			fmt.Println("criteria to game: ", crit_to_game)
			//crit_status.String()
		}
		/* ------------------------- Create a board ------------------------- */

		createNewBoard := false

		if createNewBoard {
			boardUuid := uuid.New()
			boardName := "I drink too much"

			board_status, _ := snap_sql.CreateBoard(
				boardUuid,
				readGameInGroupFromName[0].Id,
				user.Id,
				boardName,
				1)

			fmt.Println("Create Board: ", boardName, board_status)
		}
	}
}

func showUserBoard(user *data_classes.UserProfile) {
	fmt.Println()

	_, err := snap_sql.ReadUsersBoards(user.Id)

	if err != nil {
		fmt.Println("Read User Boards err: ", err)
	} else {
		fmt.Println("Read User Boards... ", user.Id)
		//fmt.Println(enc.ToIndentedJson(userBoards, "", "  "))
	}
}

func main() {

	snap_sql.SetupTables()

	createGlobalGroup()

	users := createUser()
	user := users[0]

	createBreweriesGroup(user.Id)
	createGame(user)
	showUserBoard(user)
}
