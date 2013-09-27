package caching

import (
	"fmt"
	"io/ioutil"
	"path"
)

const FilePath = "../sqlQueries/"

type CacheEntry struct {
	Path, Script string
	Err          error
}

type Entries struct {
	AddUserToGroup,
	CreateBoard,
	CreateBoardTable,
	CreateCriteria,
	CreateCriteriaTable,
	CreateGame,
	CreateGameTable,
	CreateGroup,
	CreateGroupTable,
	CreateTile,
	CreateTileTable,
	CreateUser,
	CreateUserTable,
	CreateUserToGroupTable,
	DropTable,
	ReadAllGamesInGroup,
	ReadAllUsers,
	ReadUsersBoards,
	ReadBoardFromId,
	ReadGameFromId,
	ReadGameFromName,
	ReadGameInGroupFromName,
	ReadGroup,
	ReadUserByEmail,
	ReadUserById,
	TableExists *CacheEntry
}

var CacheEntries *Entries = nil

func init() {
	CacheEntries = &Entries{
		AddUserToGroup:          provideFile("addUserToGroup"),
		CreateBoard:             provideFile("createBoard"),
		CreateBoardTable:        provideFile("createBoardTable"),
		CreateCriteria:          provideFile("createCriteria"),
		CreateCriteriaTable:     provideFile("createCriteriaTable"),
		CreateGame:              provideFile("createGame"),
		CreateGameTable:         provideFile("createGameTable"),
		CreateGroup:             provideFile("createGroup"),
		CreateGroupTable:        provideFile("createGropuTable"),
		CreateTile:              provideFile("createTile"),
		CreateTileTable:         provideFile("createTileTable"),
		CreateUser:              provideFile("createUser"),
		CreateUserTable:         provideFile("createUserTable"),
		CreateUserToGroupTable:  provideFile("createUserToGroupTable"),
		DropTable:               provideFile("drop-table"),
		ReadAllGamesInGroup:     provideFile("readAllGamesInGroup"),
		ReadAllUsers:            provideFile("readAllUsers"),
		ReadUsersBoards:         provideFile("readAllUserBoards"),
		ReadBoardFromId:         provideFile("readBoardFromId"),
		ReadGameFromId:          provideFile("readGameFromId"),
		ReadGameFromName:        provideFile("readGameFromName"),
		ReadGameInGroupFromName: provideFile("readGameInGroupFromName"),
		ReadGroup:               provideFile("readGroup"),
		ReadUserByEmail:         provideFile("readUserByEmail"),
		ReadUserById:            provideFile("readUserById"),
		TableExists:             provideFile("tableExists"),
	}
}

func provideFile(name string) *CacheEntry {
	f := path.Join(FilePath, name+".sql")
	return NewEntry(f)
}

func NewEntry(path string) (c *CacheEntry) {
	script, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		c = &CacheEntry{
			Path: path,
			Err:  err,
		}
		return c
	}

	c = &CacheEntry{
		Path:   path,
		Script: string(script),
		Err:    nil,
	}

	return c
}
