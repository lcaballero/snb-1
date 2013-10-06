package caching

import (
	"testing"
)

func Test_CacheEntries_Setup(t *testing.T) {

	hasSql(t, Cache().AddUserToGroup)
	hasSql(t, Cache().CreateBoard)
	hasSql(t, Cache().CreateBoardTable)
	hasSql(t, Cache().CreateCriteria)
	hasSql(t, Cache().CreateCriteriaTable)
	hasSql(t, Cache().CreateGame)
	hasSql(t, Cache().CreateGameTable)
	hasSql(t, Cache().CreateGroup)
	hasSql(t, Cache().CreateGroupTable)
	hasSql(t, Cache().CreateTile)
	hasSql(t, Cache().CreateTileTable)
	hasSql(t, Cache().CreateUser)
	hasSql(t, Cache().CreateUserTable)
	hasSql(t, Cache().CreateUserToGroupTable)
	hasSql(t, Cache().DropTable)
	hasSql(t, Cache().ReadAllGamesInGroup)
	hasSql(t, Cache().ReadAllUsers)
	hasSql(t, Cache().ReadUsersBoards)
	hasSql(t, Cache().ReadBoardFromId)
	hasSql(t, Cache().ReadGameFromId)
	hasSql(t, Cache().ReadGameFromName)
	hasSql(t, Cache().ReadTile)
	hasSql(t, Cache().ReadBoardTiles)
	hasSql(t, Cache().ReadGameInGroupFromName)
	hasSql(t, Cache().ReadGroup)
	hasSql(t, Cache().ReadUserByEmail)
	hasSql(t, Cache().ReadUserById)
	hasSql(t, Cache().TableExists)
	hasSql(t, Cache().CreateGameToCriteriaTable)
	hasSql(t, Cache().AddCriteriaToGame)
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
