package session_manager

type User struct {
	FirstName string
	LastName  string
	UserName  string

	InternalId string
	UserGroup  string
	Teir       string
	SessionId  string

	Email    string
	Phone    string
	Address  string
	Company  string
	Location string
	Url      string

	Created_At string
	Updated_At string
	Updated_By string
}
