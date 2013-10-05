// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Start running inserts")
// }

// func creatUser(id, version, usergroup_id int, password, username string) {
// 	sql := `
// insert into _user
// 	(id, version, usergroup_id, password, username)
// values
// 	($1, $2, $3, $4, $5);
// `
// 	result, err := getConnection().Exec(sql, id, version, usergroup_id, password, username)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// }
