package sql_utils

import (
	"database/sql"
	"fmt"
	"requests"
	"sql_utils/caching"
)

type ConnectionParameters struct {
	DataBaseName string
	ServerName   string
	UserName     string
	Password     string
	Mode         string
}

func (p *ConnectionParameters) String() string {
	return "user=lucascaballero dbname=snb password=Livebig6## sslmode=disable"
}

type Conn struct {
	db     *sql.DB
	params *ConnectionParameters
}

func (c *Conn) Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = c.db.Query(query, args...)
	c.db.Close()
	return rows, err
}

func (c *Conn) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	result, err = c.db.Exec(query, args...)
	c.db.Close()
	return result, err
}

func GetConnection() *Conn {
	params := &ConnectionParameters{}
	database, _ := sql.Open("postgres", params.String())
	return &Conn{
		db:     database,
		params: params,
	}
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

	sql := caching.Cache().TableExists.Script

	rows, err := GetConnection().Query(sql, dbName, tableName)

	if err != nil {

		fmt.Println(err)
		return false

	} else {

		m := requests.ToMapping(rows)
		return len(m) > 0
	}
}

// func DropTable(tableName string) bool {

// 	sql := "drop table if exists " + tableName + " cascade;"

// 	result, err := GetConnection().Exec(sql)

// 	if err != nil {
// 		fmt.Println("querying err: ", err)
// 		return false
// 	}

// 	rows, err1 := result.RowsAffected()

// 	if err1 != nil {
// 		fmt.Println("Error when requesting RowsAffected: ", err1)
// 	} else {
// 		fmt.Println("Dropping: "+tableName+", RowsAffected: ", rows)
// 	}

// 	return true
// }

func DropAllTables(schema string) bool {

	sql := "drop schema " + schema + " cascade;"
	sql = sql + "create schema " + schema + ";"

	result, err := GetConnection().Exec(sql)

	if err != nil {
		fmt.Println("Error dropping schemas: ", schema)
		return false
	}

	rows, err1 := result.RowsAffected()

	if err1 != nil {
		fmt.Println("Error when requesting RowsAffected: ", err1)
	} else {
		fmt.Println("Dropping: "+schema+", RowsAffected: ", rows)
	}

	return true
}
