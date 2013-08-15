package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	enc "json_helpers"
	"requests"
	"sql_text"
)

func getConnection() *sql.DB {
	database, _ := sql.Open(
		"postgres",
		"user=lucascaballero dbname=snb password=Livebig6## sslmode=disable")
	return database
}

func readAllUsers() string {
	rows, _ := getConnection().Query("SELECT * FROM _user;")
	json := toJson(rows)
	return json
}

func readUser(n int) string {
	sql := "SELECT * FROM _user WHERE id=$1"
	rows, err := getConnection().Query(sql, n)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	json := toJson(rows)
	return json
}

func creatUser(id, version, usergroup_id int, password, username string) {
	sql := `
insert into _user
	(id, version, usergroup_id, password, username)
values
	($1, $2, $3, $4, $5);
`
	result, err := getConnection().Exec(sql, id, version, usergroup_id, password, username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func createUserTable() {
	sql := `
CREATE TABLE _User (
	Id int PRIMARY KEY,
	Version int NOT NULL,

	UserGroup_Id int NOT NULL,
	Password text NOT NULL,
	UserName text NOT NULL
);
`
	result, err := getConnection().Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func tableExists(dbName, tableName string) bool {
	sql := `
SELECT *
FROM information_schema.tables
WHERE
   table_schema='public'
   and table_catalog=$1
   and table_name=$2;
`
	rows, err := getConnection().Query(sql, dbName, tableName)
	if err != nil {
		fmt.Println("Table Exists: ", err)
	}
	m := requests.ToMapping(rows)
	return len(m) == 1
}

func toJson(rows *sql.Rows) string {
	defer rows.Close()
	mapping := requests.ToMapping(rows)
	json := enc.ToIndentedJson(mapping, "", "  ")
	return json
}

func hasUserTable() bool {
	has_table := tableExists("Go_Testing", "_user")
	return has_table
}

func main() {

	if !hasUserTable() {
		fmt.Println("Creating Table")
		createUserTable()
	}

	for i := 0; i < 5; i++ {
		creatUser(i, i, 0, fmt.Sprintf("pass-%v", i), fmt.Sprintf("user-%v", i))
	}

	fmt.Println()
	fmt.Println("Get User(1)")

	json := readUser(1)
	fmt.Println()
	fmt.Println()
	fmt.Println(json)

	fmt.Println()
	fmt.Println("Get All Users")
	json = readAllUsers()
	fmt.Println()
	fmt.Println()
	fmt.Println(json)

	s := sql_text.SqlScripts{}

	fmt.Println()
	fmt.Println()
	fmt.Println(s)

	ss := sql_text.Default()

	fmt.Println()
	fmt.Println()
	fmt.Println(ss)
}
