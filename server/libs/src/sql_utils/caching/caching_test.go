package caching

import (
	"testing"
)

func Test_CacheEntries_Setup(t *testing.T) {

	LoadSqlScripts()

	hasSql(t, CacheEntries.AddUserToGroup)
	hasSql(t, CacheEntries.CreateBoard)
	hasSql(t, CacheEntries.CreateBoardTable)
	hasSql(t, CacheEntries.CreateCriteria)
	hasSql(t, CacheEntries.CreateCriteriaTable)
	hasSql(t, CacheEntries.CreateGame)
	hasSql(t, CacheEntries.CreateGameTable)
	hasSql(t, CacheEntries.CreateGroup)
	hasSql(t, CacheEntries.CreateGroupTable)
	hasSql(t, CacheEntries.CreateTile)
	hasSql(t, CacheEntries.CreateTileTable)
	hasSql(t, CacheEntries.CreateUser)
	hasSql(t, CacheEntries.CreateUserTable)
	hasSql(t, CacheEntries.CreateUserToGroupTable)
	hasSql(t, CacheEntries.DropTable)
	hasSql(t, CacheEntries.ReadAllGamesInGroup)
	hasSql(t, CacheEntries.ReadAllUsers)
	hasSql(t, CacheEntries.ReadUsersBoards)
	hasSql(t, CacheEntries.ReadBoardFromId)
	hasSql(t, CacheEntries.ReadGameFromId)
	hasSql(t, CacheEntries.ReadGameFromName)
	hasSql(t, CacheEntries.ReadTile)
	hasSql(t, CacheEntries.ReadBoardTiles)
	hasSql(t, CacheEntries.ReadGameInGroupFromName)
	hasSql(t, CacheEntries.ReadGroup)
	hasSql(t, CacheEntries.ReadUserByEmail)
	hasSql(t, CacheEntries.ReadUserById)
	hasSql(t, CacheEntries.TableExists)
}

func hasSql(t *testing.T, c *CacheEntry) {
	if c.Path == "" {
		t.Error("Cache Entry provided doesn't have a path.")
	}

	if c.Script == "" {
		t.Error("CacheEntry doesn't contain any sql code.")
	}

	if c.Err != nil {
		t.Error("An error occured while reading a CacheEntry with path: '" + c.Path + "'")
	}
}
