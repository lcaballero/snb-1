package sql_utils

import (
	"database/sql"
	"fmt"
	"requests"
	"sql_utils/caching"
)

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

	sql := caching.CacheEntries.TableExists

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
