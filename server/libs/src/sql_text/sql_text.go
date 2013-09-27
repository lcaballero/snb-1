// package sql_text

// import (
// 	"fmt"
// 	"text_files"
// )

// const (
// 	ReadAllUsers = "sql/read-all-users.sql"
// 	InsertUser = "sql/create-new-users.sql"
// 	CreateUserTable = "sql/create-user-table.sql"
// )

// type SqlScripts struct {
// 	files []string
// 	lookup map[string]string
// }

// func Default() *SqlScripts {
// 	return New([]string{
// 		ReadAllUsers,
// 		InsertUser,
// 		CreateUserTable,
// 	})
// }

// var cachedScripts *SqlScripts = nil

// func Cached() *SqlScripts {
// 	if cachedScripts == nil {
// 		cachedScripts = Default()
// 	}
// 	return cachedScripts
// }

// func New(files []string) *SqlScripts {

// 	ss := SqlScripts{ files, make(map[string]string) }

// 	for _, path := range files {
// 		text, err := text_files.ReadContent(path)

// 		if err != nil {
// 			text = ""
// 			fmt.Println("Error Reading", path)	
// 		}
// 		ss.lookup[path] = text		
// 	}

// 	return &ss;
// }

// func (s *SqlScripts) Script(key string) string {
// 	value, ok := s.lookup[key]

// 	if ok {
// 		return value
// 	} else {
// 		return ""
// 	}
// }