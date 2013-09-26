package sql_utils

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path"
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

type CacheEntry struct {
	Path, Script string
	Err          error
}

type Entries struct {
	TableExists, CreateGame *CacheEntry
}

var CacheEntries *Entries = nil

func init() {
	CacheEntries = &Entries{
		TableExists: NewEntry(path.Join(FilePath, "tableExists.sql")),
		CreateGame:  NewEntry(path.Join(FilePath, "createGame.sql")),
	}
}

func NewEntry(path string) (c *CacheEntry) {
	script, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		c = &CacheEntry{
			Path: path,
			Err:  nil,
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

	val, isString := v.(string)
	if ok && isString {
		return val
	} else {
		return ""
	}
}

func TableExists(dbName, tableName string) bool {

	sql := CacheEntries.TableExists

	if sql.Err != nil {
		fmt.Println(sql.Err)
		return false
	}

	rows, err := GetConnection().Query(sql.Script, dbName, tableName)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		m := requests.ToMapping(rows)
		hasLen := len(m) > 0

		if hasLen {
			fmt.Println("Table Exists: ", tableName)
		} else {
			fmt.Println("Table Does NOT Exists: ", tableName)
		}

		return hasLen
	}
}

func DropTable(tableName string) bool {

	sql := "drop table if exists " + tableName

	fmt.Println(sql)

	result, err := GetConnection().Exec(string(sql))

	if err != nil {
		fmt.Println("querying err: ", err)
		return false
	}

	fmt.Println("result: ", result)

	return true
}
