package snap_sql

import (
	"data_classes"
	"database/sql"
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"time"
)

// ---------------------- Read User Functions ---------------------- //

func ReadAllUsers() ([]data_classes.UserProfile, error) {

	rows, err := sql_utils.GetConnection().Query(caching.CacheEntries.ReadAllUsers.Script)

	return processUserProfiles(rows, err)
}

func ReadUserById(userId string) ([]data_classes.UserProfile, error) {

	rows, err := sql_utils.GetConnection().Query(caching.CacheEntries.ReadUserById.Script, userId)

	return processUserProfiles(rows, err)
}

func ReadUserByEmail(email string) ([]data_classes.UserProfile, error) {

	rows, err := sql_utils.GetConnection().Query(caching.CacheEntries.ReadUserByEmail.Script, email)

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
			u := data_classes.UserProfile{
				Id:        v["id"].(string),
				Email:     v["email"].(string),
				DateAdded: v["date_added"].(time.Time),
			}

			profiles[i] = u
		}

		//fmt.Println("email[0]:", profiles[0].GetProp("date_added"))
		return profiles, nil
	}
}

func HasUser(userName string) (bool, error) {
	currentUsers, err := ReadUserByEmail(userName)

	if err != nil {
		fmt.Println(err)
		return true, err // TODO: should this be true or false?
	} else if len(currentUsers) > 0 {
		return true, err
	} else {
		return false, err
	}
}

func HasUserId(userId string) (bool, error) {
	currentUsers, err := ReadUserById(userId)

	if err != nil {
		fmt.Println(err)
		return true, err // TODO: should this be true or false?
	} else if len(currentUsers) > 0 {
		return true, err
	} else {
		return false, err
	}
}
