package models

import (
	"data_classes"
	"fmt"
	_ "github.com/bmizerany/pq"
	enc "json_helpers"
	"snap_sql"
	_ "sql_utils"
	"sql_utils/caching"
	"strconv"
	"uuid"
)

const (
	usr                = "d333"
	pw                 = "ro"
	global_group_name  = "global_group"
	global_group_desc  = "group that contains every user"
	brewery_group_name = "Breweries"
	brewery_group_desc = "Breweries in Boulder"
	game_name          = "Boulder Breweries"
	game_desc          = "Have a brew with a brewer"
)

type InitialModel struct{}

func createUser() []*data_classes.UserProfile {

	status, _ := snap_sql.CreateUser(usr, pw)

	fmt.Println("Create User: ", status)

	myUser, _ := snap_sql.ReadUserByEmail(usr)

	return myUser
}

func createGlobalGroup() {

	hasGlobalGroup, _ := snap_sql.HasGroup(global_group_name)

	if !hasGlobalGroup {

		globalGroupUuid := uuid.New()
		group_status, _ := snap_sql.CreateGroup(
			globalGroupUuid,
			global_group_name,
			global_group_desc,
			globalGroupUuid)

		fmt.Println("Create Group: ", group_status)

	} else {

		fmt.Println("Has Group: ", global_group_name)
	}
}

func createBreweriesGroup(userId string) {

	hasBreweryGroup, _ := snap_sql.HasGroup(brewery_group_name)

	if !hasBreweryGroup {

		groupUuid := uuid.New()
		group_status, _ := snap_sql.CreateGroup(
			groupUuid,
			brewery_group_name,
			brewery_group_desc,
			userId)

		fmt.Println("Create Group: ", brewery_group_name, group_status)

	} else {

		fmt.Println("Has Group: ", brewery_group_name)
	}
}

func createGame(group data_classes.GroupData) {

	gameUuid := uuid.New()

	createGame_status, _ := snap_sql.CreateGame(
		gameUuid,
		group.Id,
		game_name,
		game_desc)

	fmt.Println("Create Game: ", game_name, createGame_status)
	fmt.Println("Read Game from Id: ...")

	readGameFromId, _ := snap_sql.ReadGameFromId(gameUuid)

	fmt.Println(enc.ToIndentedJson(readGameFromId, "", "  "))
}

func createCriteria(group data_classes.GameData) {

	for i := 0; i < 25; i++ {

		criteriaUuid := uuid.New()

		crit_status, err := snap_sql.CreateCriteria(
			criteriaUuid,
			"crit_"+strconv.Itoa(i),
		)

		if err != nil {
			fmt.Println(err, crit_status)
		}

		critToGameUuid := uuid.New()

		crit_to_game, err := snap_sql.AddCriteriaToGame(
			critToGameUuid,
			group.Id,
			criteriaUuid,
			1,
			1)

		fmt.Println("criteria to game: ", crit_to_game)
	}
}

func createBoard(user *data_classes.UserProfile, game data_classes.GameData) {

	boardUuid := uuid.New()
	boardName := "I drink too much"

	board_status, _ := snap_sql.CreateBoard(
		boardUuid,
		game.Id,
		user.Id,
		boardName,
		1)

	fmt.Println("Create Board: ", boardName, board_status)
}

func createFullGame(user *data_classes.UserProfile) {

	hasBreweryGroup, _ := snap_sql.HasGroup(brewery_group_name)

	if hasBreweryGroup {

		breweryGroup, _ := snap_sql.ReadGroup(brewery_group_name)

		createGame(breweryGroup[0])

		fmt.Println()
		fmt.Println("Read Game in group from Name: ...")

		game, _ := snap_sql.ReadGameInGroupFromName(
			breweryGroup[0].Id,
			game_name)

		createCriteria(game[0])
		createBoard(user, game[0])
	}
}

func showUserBoard(user *data_classes.UserProfile) {

	_, err := snap_sql.ReadUsersBoards(user.Id)

	if err != nil {

		fmt.Println("Read User Boards err: ", err)

	} else {

		fmt.Println("Read User Boards... ", user.Id)
	}
}

func (m *InitialModel) Initialize() {

	caching.LoadSqlScripts()
	snap_sql.DropAllTables()
	snap_sql.SetupTables()

	createGlobalGroup()

	users := createUser()
	user := users[0]

	createBreweriesGroup(user.Id)
	createFullGame(user)
	showUserBoard(user)
}
