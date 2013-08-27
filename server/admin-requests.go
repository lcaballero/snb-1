package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	enc "json_helpers"
	"requests"
	//"sql_text"
)

func getConnection() *sql.DB {
	database, _ := sql.Open(
		"postgres",
		"user=lucascaballero dbname=snb password=Livebig6## sslmode=disable")
	return database
}

func ReadAllUsers() string {
	rows, _ := getConnection().Query("SELECT * FROM _user;")
	json := toJson(rows)
	return json
}

func ReadUser(n string) string {
	sql := "SELECT * FROM _user WHERE id=$1"
	rows, err := getConnection().Query(sql, n)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	json := toJson(rows)
	return json
}

func ReadUserFromEmail(n string) string {
	sql := "SELECT * FROM _user WHERE email=$1"
	rows, err := getConnection().Query(sql, n)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	json := toJson(rows)
	return json
}

func CreatUser(email, password string) {
	sql := `
insert into _user
	(email, password, date_added)
values
	($1, $2, now());
`
	result, err := getConnection().Exec(sql, email, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func CreateGroup(group_name, group_desc string) {
	sql := `
	insert into SocialGroup
		(group_name, group_desc, date_added)
	values
		($1, $2, now());
	`

	result, err := getConnection().Exec(sql, group_name, group_desc)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
/*
func createUserTable() {
	sql := `
CREATE TABLE _User (
	id bigserial NOT NULL,
	Email varchar(255),
	Password text NOT NULL,
	date_added timestamp without time zone,
	CONSTRAINT _user_pkey PRIMARY KEY (id)
);
`
	result, err := getConnection().Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func createGroupsTable() {
	sql := `
CREATE TABLE SocialGroup
(
  id bigserial NOT NULL,
  group_name character varying(40) NOT NULL,
  group_desc text NOT NULL,
  date_added timestamp without time zone,
  CONSTRAINT social_group_pkey PRIMARY KEY (id)
);
`
	result, err := getConnection().Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func createUserToGroupTable() {
	sql := `
CREATE TABLE UserToGroup
(
  id bigserial NOT NULL,
  group_id bigserial NOT NULL,
  user_id bigserial NOT NULL,
  date_added timestamp without time zone,
  CONSTRAINT user_to_group_pkey PRIMARY KEY (id)
);
`
	result, err := getConnection().Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
*/


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

func hasGroupTable() bool {
	has_table := tableExists("Go_Testing", "SocialGroup")
	return has_table
}

func hasUserToGroupTable() bool {
	has_table := tableExists("Go_Testing", "UserToGroup")
	return has_table
}
