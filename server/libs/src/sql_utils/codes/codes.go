package codes

// ---------------------- Status Codes ---------------------- //

type StatusCode struct {
	Code int
	Msg  string
}

const (
	success               = 200
	user_exists           = 101
	user_does_not_exists  = 102
	group_exists          = 103
	group_does_not_exists = 104
	file_read_error       = 501
	db_error              = 502
)

// var STATUS_CODES = map[int]StatusCode{
// 	SUCCESS:               {SUCCESS, "Success"},
// 	USER_EXISTS:           {USER_EXISTS, "User already exists"},
// 	USER_DOES_NOT_EXISTS:  {USER_DOES_NOT_EXISTS, "User Does Not Exists"},
// 	GROUP_EXISTS:          {GROUP_EXISTS, "Group already exists"},
// 	GROUP_DOES_NOT_EXISTS: {GROUP_DOES_NOT_EXISTS, "Group Does Not Exists"},
// 	FILE_READ_ERR:         {FILE_READ_ERR, "File Read Error"},
// 	db_error:                {db_error, "Database Error"},
// }

var (
	Success                = StatusCode{success, "Sucess"}
	User_Exists            = StatusCode{user_exists, "User already exists."}
	User_Does_Not_Exists   = StatusCode{user_does_not_exists, "User does not exists."}
	Group_Exists           = StatusCode{group_exists, "Group exists"}
	Group_Does_Not_Exsists = StatusCode{group_does_not_exists, "Group does not exists."}
	File_Read_Error        = StatusCode{file_read_error, "File read error."}
	Db_Error               = StatusCode{db_error, "Data base error."}
)
