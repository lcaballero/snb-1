package sql_utils

import (
	"database/sql"
	"requests"
)

const FilePath = "../sqlQueries/"

// ---------------------- Status Codes ---------------------- //

type StatusCode struct {
	Code int
	Msg  string
}

const (
	SUCCESS               = 200
	USER_EXISTS           = 101
	USER_DOES_NOT_EXISTS  = 102
	GROUP_EXISTS          = 103
	GROUP_DOES_NOT_EXISTS = 104
	FILE_READ_ERR         = 501
	DB_ERR                = 502
)

var STATUS_CODES = map[int]StatusCode{
	SUCCESS:               {SUCCESS, "Success"},
	USER_EXISTS:           {USER_EXISTS, "User already exists"},
	USER_DOES_NOT_EXISTS:  {USER_DOES_NOT_EXISTS, "User Does Not Exists"},
	GROUP_EXISTS:          {GROUP_EXISTS, "Group already exists"},
	GROUP_DOES_NOT_EXISTS: {GROUP_DOES_NOT_EXISTS, "Group Does Not Exists"},
	FILE_READ_ERR:         {FILE_READ_ERR, "File Read Error"},
	DB_ERR:                {DB_ERR, "Database Error"},
}

func GetConnection() *sql.DB {
	database, _ := sql.Open(
		"postgres",
		"user=lucascaballero dbname=snb password=Livebig6## sslmode=disable")
	return database
}

func ToSqlMap(rows *sql.Rows) []map[string]interface{} {
	defer rows.Close()
	return requests.ToMapping(rows)
}

func ObjToString(ref string, o map[string]interface{}) string {
	v, ok := o[ref]

	if ok && v != nil {
		return v.(string)
	} else {
		return ""
	}
}
