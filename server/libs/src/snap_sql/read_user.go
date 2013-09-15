package snap_sql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"data_classes"
	"sql_utils"
)


const FilePath = "../sqlQueries/"

// ---------------------- Read User Functions ---------------------- //

func ReadAllUsers() ([]data_classes.UserProfile, error) {
	// "SELECT * FROM _user;"
	sql, err := ioutil.ReadFile(FilePath + "readAllUsers.sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := sql_utils.GetConnection().Query(string(sql))

	return processUserProfiles(rows, err)
}

func ReadUserById(userId string) ([]data_classes.UserProfile, error) {
	//sql := "SELECT * FROM _user WHERE id=$1"
	sql, err := ioutil.ReadFile(FilePath + "readUserById.sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := sql_utils.GetConnection().Query(string(sql), userId)
	
	return processUserProfiles(rows, err)
}

func ReadUserByEmail(email string) ([]data_classes.UserProfile, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql, err := ioutil.ReadFile(FilePath + "readUserByEmail.sql");

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := sql_utils.GetConnection().Query(string(sql), email)

	return processUserProfiles(rows, err)
}

func processUserProfiles(sqlRows *sql.Rows, err error) ([]data_classes.UserProfile, error) {
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		mappedRows := sql_utils.ToSqlMap(sqlRows)

		profiles := make([]data_classes.UserProfile, len(mappedRows))

		for i, v := range mappedRows {
			//fmt.Println(enc.ToIndentedJson(v, "", "  "))
			anchor := data_classes.Anchor{}
			anchor.SetMap(v)
			profiles[i] = data_classes.UserProfile{Anchor:anchor}
		}

		//fmt.Println("email[0]:", profiles[0].GetProp("date_added"))
		return profiles, nil
	}
}
