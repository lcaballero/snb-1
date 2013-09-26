package codes

import (
	"uuid"
)

type StatusCode struct {
	code    int
	msg     string
	session string
	salt    string
}

const (
	salt                  = "Saltiness"
	success               = 200
	user_exists           = 101
	user_does_not_exists  = 102
	group_exists          = 103
	group_does_not_exists = 104
	file_read_error       = 501
	db_error              = 502
)

var (
	Success                = newCode(success, "Sucess")
	User_Exists            = newCode(user_exists, "User already exists.")
	User_Does_Not_Exists   = newCode(user_does_not_exists, "User does not exists.")
	Group_Exists           = newCode(group_exists, "Group exists")
	Group_Does_Not_Exsists = newCode(group_does_not_exists, "Group does not exists.")
	File_Read_Error        = newCode(file_read_error, "File read error.")
	Db_Error               = newCode(db_error, "Data base error.")
)

func newCode(code int, msg string) StatusCode {
	id := uuid.New()
	c := StatusCode{
		code,
		msg,
		id,
		salt,
	}

	return c
}
