package main

import (
	"snap_sql"
	"sql_utils"
	"sql_utils/caching"
	"testing"
)

const (
	email_1 = "user-1@host.com"
	pass_1  = "user-1"
)

func init() {
	caching.LoadSqlScripts()
	sql_utils.DropTable("_user")
	snap_sql.CreateUserTable()
	snap_sql.CreateUser(email_1, pass_1)
}

func Test_Can_Read_Users(t *testing.T) {
	users, _ := ReadAllUsers()

	if len(users) <= 0 {
		t.Error("No users read from DB")
	}

	if len(users) > 1 {
		t.Error("Shouldn't have found more than 1 user.  Found: ", len(users))
	}

	u := users[0]

	if u.Email != email_1 {
		t.Error("Didn't load the right email -- or could find the 1 user.", "Found email: ", "'"+u.Email+"'", users[0])
	}
}
