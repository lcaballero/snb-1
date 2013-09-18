package snap_sql


import(
	"fmt"
	"io/ioutil"
	"sql_utils"
)

func CreateGame(gameUuid, groupId, name, description string) (sql_utils.StatusCode, error) {

	var status sql_utils.StatusCode;

	sql, err := ioutil.ReadFile(sql_utils.FilePath+"createGame.sql")

	if err != nil {
		fmt.Println(err)
		status = sql_utils.STATUS_CODES[sql_utils.FILE_READ_ERR]
	} else {
		//gameUuid := uuid.New()

		// ?? should we assume groupId is valid or should 
		// we run a db query to ensure it's valid?

		// winning_board_id == null
		// active is set to false to begin with until at least 25 criteria
		// are associated with the board
		_, err := sql_utils.GetConnection().Exec(
			string(sql), 
			gameUuid, groupId, name, description, 1)
		
		if err != nil {
			fmt.Println(err)
			status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
		} else {
			status = sql_utils.STATUS_CODES[sql_utils.SUCCESS]
			return status, err
		}
	}


	return status, err
	
}