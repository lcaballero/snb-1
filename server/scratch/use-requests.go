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
	//"sql_text"
)

const FilePath = "../sqlQueries/"

func getConnection() *sql.DB {
	database, _ := sql.Open(
		"postgres",
		"user=lucascaballero dbname=snb password=Livebig6## sslmode=disable")
	return database
}

/* ---------------------- Anchor Dictionary..ish ---------------------- */

type Anchor struct {
	refMap map[string]interface{}
}

func (anchor Anchor) SetMap(m map[string]interface{}){
	anchor.refMap = m;
}

func (anchor Anchor) GetProp(reqField string) string {
	oField, ok := anchor.refMap[reqField]

	if !ok {
		fmt.Println("refMap err: ", ok)
		return ""
	}

	// TODO: need to abstract the type cast so we can use
	// other type such as int.
	field, ok := oField.(string)
	
	if ok {
		return field
	} else {
		fmt.Println("Error: Unable to convert %v to a string", reqField)
		return ""
	}
}

/* ---------------------- User Profile ---------------------- */

type Printer interface {
	PrintAll() map[string]interface{}
}

type UserProfile struct {
	Anchor
}

func (userProfile UserProfile) PrintAll() map[string]interface{} {
	m := make(map[string]interface{})

	m["Email"] = userProfile.Email()
	m["Date_Added"] = userProfile.DateAdded()

	return m
}

func (userProfile UserProfile) Email() string {
	return userProfile.GetProp("email")
}

func (userProfile UserProfile) DateAdded() string {
	return userProfile.GetProp("date_added")
}

// ---------------------- Read User Functions ---------------------- //

func readAllUsers() ([]UserProfile, error) {
	// "SELECT * FROM _user;"
	sql, err := ioutil.ReadFile(FilePath + "readAllUsers.sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := getConnection().Query(string(sql))

	return processUserProfiles(rows, err)
}

func readUserById(n int) ([]UserProfile, error) {
	//sql := "SELECT * FROM _user WHERE id=$1"
	sql, err := ioutil.ReadFile(FilePath + "readUserById.sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := getConnection().Query(string(sql), n)
	
	return processUserProfiles(rows, err)
}

func readUserByEmail(email string) ([]UserProfile, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql, err := ioutil.ReadFile(FilePath + "readUserByEmail.sql");

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := getConnection().Query(string(sql), email)

	return processUserProfiles(rows, err)
}

func processUserProfiles(sqlRows *sql.Rows, err error) ([]UserProfile, error) {
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		mappedRows := toSqlMap(sqlRows)

		profiles := make([]UserProfile, len(mappedRows))

		for i, v := range mappedRows {
			//fmt.Println(enc.ToIndentedJson(v, "", "  "))
			profiles[i] = UserProfile{Anchor:Anchor{refMap:v}}
		}

		//fmt.Println("email[0]:", profiles[0].GetProp("date_added"))
		return profiles, nil
	}
}

// ---------------------- User Functions ---------------------- //

// func userExists(email) {
// 	json := readUserByEmail(email)
// }

func createUser(email, password string) (bool, error) {

	has_user, err := hasUser(email)

	if err != nil {
		fmt.Println(err)
	} else if has_user {
		fmt.Println("User already exists in DB _user")
	} else {

		sql, err := ioutil.ReadFile(FilePath+"createUser.sql")

		if err != nil {
			fmt.Println(err)
		} else {
			myUuid := uuid.New()
			result, err := getConnection().Exec(string(sql), myUuid, email, password)
			
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Create User result: ", result)
				return true, err
			}
		
		}
	}
	return false, err
}


func createUserTable() {

	sql, err := ioutil.ReadFile(FilePath + "createUserTable.sql")

	if err != nil {
		fmt.Println(err)
	} else {
		result, err := getConnection().Exec(string(sql))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
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

func createGroup(group_name, group_desc string) {
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

func tableExists(dbName, tableName string) bool {

	sql, err := ioutil.ReadFile(FilePath + "tableExists.sql")

	if err != nil {
		fmt.Println(err)
		return false
	}

	rows, err := getConnection().Query(string(sql), dbName, tableName)
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
	has_table := tableExists("snb", "SocialGroup")
	return has_table
}

func hasUserToGroupTable() bool {
	has_table := tableExists("snb", "UserToGroup")
	return has_table
}

func hasUser(userName string) (bool, error) {
	currentUsers, err := readUserByEmail(userName)

	if err != nil {
		fmt.Println(err)
		return true, err // TODO: should this be true or false?
	} else if len(currentUsers) > 0 {
		return true, err
	} else {
		return false, err
	}
}
/* ------------------------- Main ------------------------- */

func main() {

	if !hasUserTable() {
		fmt.Println("Creating Table")
		createUserTable()
	}
/*
	if !hasGroupTable() {
		fmt.Println("Create Group table")
		createGroupsTable()
	}

	if !hasUserToGroupTable() {
		fmt.Println("Create UserToGroup table")
		createUserToGroupTable()
	}
*/
	/* ------------------------- Create User ------------------------- */

	usr := "Lucas"
	pw := "ro"

	was_created, err := createUser(usr, pw)

	if was_created {
		fmt.Println("User was created: ", usr)
	} 

	/* ------------------------- Read User By Email ------------------------- */
	userByEmail, err := readUserByEmail("Ryan")

	fmt.Println()
	fmt.Println()

	for i := 0; i < len(userByEmail); i++ {
		fmt.Println(enc.ToIndentedJson(userByEmail[i].PrintAll(), "", "  "));	
	}
	
	/* ------------------------- Read All Users ------------------------- */
	allUsers, err := readAllUsers()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println()

	for i := 0; i < len(allUsers); i++ {
		fmt.Println(enc.ToIndentedJson(allUsers[i].PrintAll(), "", "  "));	
	}

	/*
	createGroup("global_group", "group that contains every user")

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
