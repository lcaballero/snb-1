package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"requests"
	enc "json_helpers"
	_ "github.com/lib/pq"
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

func toJson(rows *sql.Rows) string {
	defer rows.Close()
	mapping := requests.ToMapping(rows)
	json := enc.ToIndentedJson(mapping, "", "  ")
	return json
}

func getAllUser(w http.ResponseWriter, r *http.Request){
	h := w.Header()
	h["Content-Type"] = []string { "application/json", "charset-utf-8", }	

	users := readAllUsers()

	fmt.Fprint(w, users)
}