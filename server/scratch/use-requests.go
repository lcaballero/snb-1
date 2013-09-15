package main

import (
	"database/sql"
	"fmt"
	//"builtin"
	_ "github.com/bmizerany/pq"
	enc "json_helpers"
	"requests"
	"io/ioutil"
	"uuid"
	"time"
	"data_classes"
	"sql_utils"
	"snap_sql"
	//"reflect"
	//"sql_text"
)

/*
const FilePath = "../sqlQueries/"

func getConnection() *sql.DB {
	database, _ := sql.Open(
		"postgres",
		"user=lucascaballero dbname=snb password=Livebig6## sslmode=disable")
	return database
}
*/


/* ---------------------- Group Data ---------------------- */

type GroupData struct {
	data_classes.Anchor
}

func (group GroupData) GroupId() string {
	//fmt.Println("--------", group.GetProp("id").(string))
	return group.GetProp("id").(string)
}

func (group GroupData) GroupName() string {
	return group.GetProp("group_name").(string)
}

func (group GroupData) Description() string {
	return group.GetProp("group_desc").(string)
}

func (group GroupData) DateAdded() time.Time {
	return group.GetProp("date_added").(time.Time)
}





func processGroup(sqlRows *sql.Rows, err error) ([]GroupData, error) {
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		mappedRows := toSqlMap(sqlRows)

		groups := make([]GroupData, len(mappedRows))

		for i, v := range mappedRows {
			anchor := data_classes.Anchor{}
			anchor.SetMap(v)
			groups[i] = GroupData{Anchor:anchor}
		}

		//fmt.Println("group[0]:", groups[0].GetProp("group_name"))
		return groups, nil
	}
}


// ---------------------- Read Group Functions ---------------------- //

func readGroup(group_name string) ([]GroupData, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "readGroup.sql");

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		rows, err := sql_utils.GetConnection().Query(string(sql), group_name)

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			return processGroup(rows, err)
		}
	}
}

// ---------------------- Status Codes ---------------------- //

type StatusCode struct {
	Code int
	Msg string
}

var SUCCESS = 200
var USER_EXISTS = 101
var USER_DOES_NOT_EXISTS = 102
var GROUP_EXISTS = 103
var GROUP_DOES_NOT_EXISTS = 104
var FILE_READ_ERR = 501
var DB_ERR = 502

var STATUS_CODES = map[int]StatusCode {
	SUCCESS: {SUCCESS, "Success"},
	USER_EXISTS: {USER_EXISTS, "User already exists"},
	USER_DOES_NOT_EXISTS: {USER_DOES_NOT_EXISTS, "User Does Not Exists"},
	GROUP_EXISTS: {GROUP_EXISTS, "Group already exists"},
	GROUP_DOES_NOT_EXISTS: {GROUP_DOES_NOT_EXISTS, "Group Does Not Exists"},
	FILE_READ_ERR: {FILE_READ_ERR, "File Read Error"},
	DB_ERR: {DB_ERR, "Database Error"},
}

// ---------------------- User Functions ---------------------- //

func createUser(email, password string) (StatusCode, error) {

	has_user, err := hasUser(email)

	var status StatusCode;

	if err != nil {
		fmt.Println(err)
		status = STATUS_CODES[DB_ERR]
	} else if has_user {
		status = STATUS_CODES[USER_EXISTS]
	} else {

		sql, err := ioutil.ReadFile(sql_utils.FilePath+"createUser.sql")

		if err != nil {
			fmt.Println(1, err)
			status = STATUS_CODES[FILE_READ_ERR]
		} else {
			userUuid := uuid.New()
			_, err := sql_utils.GetConnection().Exec(string(sql), userUuid, email, password)
			
			if err != nil {
				fmt.Println(2, err)
				status = STATUS_CODES[DB_ERR]
			} else {
				//fmt.Println("Create User result: ", result)

				group, err := readGroup("global_group")

				if err != nil {
					status = STATUS_CODES[DB_ERR]
				} else {

					status, err := addUserToGroup(userUuid, group[0].GroupId())
					fmt.Println(status.Msg)

					if err != nil{
						fmt.Println(err)
						
					} else {
						status = STATUS_CODES[SUCCESS]
						return status, err
					}
				}
			}
		
		}
	}

	return status, err
}

func addUserToGroup(userId, groupId string) (StatusCode, error) {

	var status StatusCode;
	// has_group, err := hasGroup(groupId)

	// if has_group {

		// has_user, err := hasUserId(userId)

		//if has_user && err == nil {
			sql, err := ioutil.ReadFile(sql_utils.FilePath + "addUserToGroup.sql")

			if err != nil {
				fmt.Println(err)
				status = STATUS_CODES[DB_ERR]
			} else {

				// add user to the global group
				rowUuid := uuid.New()
				_, err := sql_utils.GetConnection().Exec(string(sql), rowUuid, groupId, userId)
fmt.Println(3)
				if err != nil {
					fmt.Println(err)
					status = STATUS_CODES[DB_ERR]
				} else {
					status = STATUS_CODES[SUCCESS]
					return status, err
				}
			}
		// } else {
		// 	status = STATUS_CODES[USER_DOES_NOT_EXISTS]
		// }
	// } else {
	// 	status = STATUS_CODES[GROUP_DOES_NOT_EXISTS]
	// }

	fmt.Println(err)
	return status, err
}

func createUserTable() {

	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createUserTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		result, err := sql_utils.GetConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}

// ---------------------- Group Functions ---------------------- //

func createGroupsTable() {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createGroupsTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		result, err := sql_utils.GetConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}

func createUserToGroupTable() {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createUserToGroupTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		result, err := sql_utils.GetConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}

func createGroup(group_name, group_desc, group_owner string) (StatusCode, error) {
		
	has_group, err := hasGroup(group_name)

	var status StatusCode;

	if err != nil {
		fmt.Println(err)
		status = STATUS_CODES[DB_ERR]
	} else if has_group {
		status = STATUS_CODES[GROUP_EXISTS]
	} else {
		sql, err := ioutil.ReadFile(sql_utils.FilePath + "createGroup.sql")

		if err != nil {
			fmt.Println(err)
			status = STATUS_CODES[FILE_READ_ERR]
		} else {
			groupUuid := uuid.New()
			_, err := sql_utils.GetConnection().Exec(string(sql), groupUuid, group_name, group_desc, group_owner)

			if err != nil {
				fmt.Println(err)
				status = STATUS_CODES[DB_ERR]
			} else {
				status = STATUS_CODES[SUCCESS]
				return status, err
			}
		}
	}

	return status, err
}

func tableExists(dbName, tableName string) bool {

	sql, err := ioutil.ReadFile(sql_utils.FilePath + "tableExists.sql")

	if err != nil {
		fmt.Println(err)
		return false
	}

	rows, err := sql_utils.GetConnection().Query(string(sql), dbName, tableName)
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

func toSqlMap(rows *sql.Rows) []map[string]interface{} {
	defer rows.Close()
	return requests.ToMapping(rows)
}

func toJson(rows *sql.Rows) string {
	defer rows.Close()
	mapping := requests.ToMapping(rows)
	json := enc.ToIndentedJson(mapping, "", "  ")
	return json
}

/* ------------------------- Has methods ------------------------- */

func hasUserTable() bool {
	has_table := tableExists("snb", "_user")
	return has_table
}

func hasGroupTable() bool {
	has_table := tableExists("snb", "socialgroup")
	return has_table
}

func hasUserToGroupTable() bool {
	has_table := tableExists("snb", "usertogroup")
	return has_table
}

func hasUser(userName string) (bool, error) {
	currentUsers, err := snap_sql.ReadUserByEmail(userName)

	if err != nil {
		fmt.Println(err)
		return true, err // TODO: should this be true or false?
	} else if len(currentUsers) > 0 {
		return true, err
	} else {
		return false, err
	}
}

func hasUserId(userId string) (bool, error) {
	currentUsers, err := snap_sql.ReadUserById(userId)

	if err != nil {
		fmt.Println(err)
		return true, err // TODO: should this be true or false?
	} else if len(currentUsers) > 0 {
		return true, err
	} else {
		return false, err
	}
}

func hasGroup(groupName string) (bool, error) {
	groups, err := readGroup(groupName)

	if err != nil {
		fmt.Println(err)
		return true, err // TODO: should this be true or false?
	} else if len(groups) > 0 {
		return true, err
	} else {
		return false, err
	}
}


/* ------------------------- Main ------------------------- */

func main() {

	if !hasUserTable() {
		fmt.Println("Creating User Table")
		createUserTable()
	}

	if !hasGroupTable() {
		fmt.Println("Create Group table")
		createGroupsTable()
	}

	if !hasUserToGroupTable() {
		fmt.Println("Create UserToGroup table")
		createUserToGroupTable()
	}

	/* ------------------------- Create Group ------------------------- */

	hasGlobalGroup, err := hasGroup("global_group")
	if !hasGlobalGroup {
		globalGroupUuid := uuid.New()
		group_status, _ := createGroup("global_group", "group that contains every user", globalGroupUuid)
		fmt.Println("Create Group: ", group_status.Msg)
	}

	/* ------------------------- Create User ------------------------- */

	usr := "d333"
	pw := "ro"

	status, err := createUser(usr, pw)

	fmt.Println("Create User: ", status.Msg)

	/* ------------------------- Read User By Email ------------------------- */
	//userByEmail, err := readUserByEmail("Ryan")

	fmt.Println()
	fmt.Println()

	// for i := 0; i < len(userByEmail); i++ {
	// 	fmt.Println(enc.ToIndentedJson(userByEmail[i].PrintAll(), "", "  "));	
	// }
	
	/* ------------------------- Read All Users ------------------------- */
	allUsers, err := snap_sql.ReadAllUsers()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println()

	for i := 0; i < len(allUsers); i++ {
		fmt.Println(enc.ToIndentedJson(allUsers[i].PrintAll(), "", "  "));	
	}

	/*

	for i := 0; i < 5; i++ {
		creatUser(fmt.Sprintf("email-%v", i), fmt.Sprintf("pass-%v", i))
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
	*/
}
