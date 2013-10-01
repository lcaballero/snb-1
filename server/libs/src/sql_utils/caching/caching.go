package caching

import (
	//	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const DefaultFilePath = "../sqlQueries/"

var internalFilePath string = DefaultFilePath

type PathProvider func(string) string

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
	ReadCriteriaFromId,
	ReadGameFromId,
	ReadGameFromName,
	ReadGameInGroupFromName,
	ReadGroup,
	ReadUserByEmail,
	ReadUserById,
	TableExists *CacheEntry
}

var CacheEntries *Entries = nil
var SqlPathProvider PathProvider = nil

func init() {
	val := FindFlag("--config-file=", os.Args)
	if val != "" {
		internalFilePath = val
		LoadSqlScripts()
	}
}

func FindFlag(flag string, args []string) string {

	val := ""

	for _, e := range args {
		hasPrefix := strings.HasPrefix(e, flag)
		if hasPrefix {
			val = e[len(flag):]
		}
	}

	return val
}

func LoadSqlScripts() {
	CacheEntries = &Entries{
		AddUserToGroup:          provideFile("addUserToGroup"),
		CreateBoard:             provideFile("createBoard"),
		CreateBoardTable:        provideFile("createBoardTable"),
		CreateCriteria:          provideFile("createCriteria"),
		CreateCriteriaTable:     provideFile("createCriteriaTable"),
		CreateGame:              provideFile("createGame"),
		CreateGameTable:         provideFile("createGameTable"),
		CreateGroup:             provideFile("createGroup"),
		CreateGroupTable:        provideFile("createGroupsTable"),
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
		ReadCriteriaFromId:      provideFile("readCriterial"),
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

	f := ""

	if SqlPathProvider == nil {
		f = path.Join(internalFilePath, name+".sql")
	} else {
		f = SqlPathProvider(name)
	}

	return NewEntry(f)
}

func NewEntry(path string) (c *CacheEntry) {

	script, err := ioutil.ReadFile(path)

	if err != nil {
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
