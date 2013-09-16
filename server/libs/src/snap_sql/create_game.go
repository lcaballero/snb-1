package snap_sql

/*
import(
	"fmt"
	"uuid"
	"io/ioutil"
	"sql_utils"
)*/

func CreateGame() bool {//id, groupId, name, description, sponsorId string, active bool) (sql_utils.StatusCode, error) {

	/*
	var status sql_utils.StatusCode;

	sql, err := ioutil.ReadFile(sql_utils.FilePath+"createGame.sql")

	if err != nil {
		fmt.Println(1, err)
		status = sql_utils.STATUS_CODES[sql_utils.FILE_READ_ERR]
	} else {
		userUuid := uuid.New()
		_, err := sql_utils.GetConnection().Exec(string(sql), userUuid, groupId, null, name, description, sponsorId, true)
		
		if err != nil {
			fmt.Println(2, err)
			status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
		} else {
			//fmt.Println("Create User result: ", result)

			group, err := ReadGroup("global_group")

			if err != nil {
				status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
			} else {

				status, err := AddUserToGroup(userUuid, group[0].Id())
				fmt.Println(status.Msg)

				if err != nil{
					fmt.Println(err)
					
				} else {
					status = sql_utils.STATUS_CODES[sql_utils.SUCCESS]
					return status, err
				}
			}
		}
	}


	return status, err
	*/
	return false
}