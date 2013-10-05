package caching

import (
	"fmt"
	"io/ioutil"
	"path"
	"rt_config"
)

const (
	DefaultFilePath = "../sqlQueries/"
)

var (
	internalFilePath string       = DefaultFilePath
	entries          *Entries     = nil
	SqlPathProvider  PathProvider = nil
)

type PathProvider func(string) string

type CacheEntry struct {
	Path, Script string
	Err          error
}

type Entries struct {
	AddUserToGroup,
	AddCriteriaToGame,
	CreateBoard,
	CreateBoardTable,
	CreateCriteria,
	CreateCriteriaTable,
	CreateGame,
	CreateGameTable,
	CreateGameToCriteriaTable,
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
	ReadBoardTiles,
	ReadCriteriaFromId,
	ReadGameFromId,
	ReadGameFromName,
	ReadGameInGroupFromName,
	ReadGroup,
	ReadTile,
	ReadUserByEmail,
	ReadUserById,
	TableExists *CacheEntry
}

func Cache() *Entries {

	return &Entries{
		AddUserToGroup:            provideFile("addUserToGroup"),
		AddCriteriaToGame:         provideFile("addCriteriaToGame"),
		CreateBoard:               provideFile("createBoard"),
		CreateBoardTable:          provideFile("createBoardTable"),
		CreateCriteria:            provideFile("createCriteria"),
		CreateCriteriaTable:       provideFile("createCriteriaTable"),
		CreateGame:                provideFile("createGame"),
		CreateGameTable:           provideFile("createGameTable"),
		CreateGameToCriteriaTable: provideFile("createGameToCriteriaTable"),
		CreateGroup:               provideFile("createGroup"),
		CreateGroupTable:          provideFile("createGroupsTable"),
		CreateTile:                provideFile("createTile"),
		CreateTileTable:           provideFile("createTileTable"),
		CreateUser:                provideFile("createUser"),
		CreateUserTable:           provideFile("createUserTable"),
		CreateUserToGroupTable:    provideFile("createUserToGroupTable"),
		DropTable:                 provideFile("drop-table"),
		ReadAllGamesInGroup:       provideFile("readAllGamesInGroup"),
		ReadAllUsers:              provideFile("readAllUsers"),
		ReadUsersBoards:           provideFile("readAllUserBoards"),
		ReadBoardFromId:           provideFile("readBoardFromId"),
		ReadBoardTiles:            provideFile("readBoardTiles"),
		ReadCriteriaFromId:        provideFile("readCriterial"),
		ReadGameFromId:            provideFile("readGameFromId"),
		ReadGameFromName:          provideFile("readGameFromName"),
		ReadGameInGroupFromName:   provideFile("readGameInGroupFromName"),
		ReadGroup:                 provideFile("readGroup"),
		ReadTile:                  provideFile("ReadTile"),
		ReadUserByEmail:           provideFile("readUserByEmail"),
		ReadUserById:              provideFile("readUserById"),
		TableExists:               provideFile("tableExists"),
	}
}

func init() {

	conf := rt_config.LoadFromCommandLine()
	env := conf.Dev

	fmt.Println("Config file:", conf.ConfigFile)
	dir := path.Dir(conf.ConfigFile)

	SqlPathProvider = func(name string) string {
		return path.Join(dir, env.SqlScripts, name+".sql")
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
