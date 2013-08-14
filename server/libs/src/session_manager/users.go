package session_manager

import (
	"fmt"
	"session_manager/hashes"
)

func AddUser(u *User) (uu *User, n interface{}, err error) {

	if u.SessionId == "" {
		fmt.Println("User's session Id should not be empty", u)
		return u, nil, nil
	}

	n, err = GetConn().Do(hashes.HMSET,
		u.InternalId,
		"FirstName", u.FirstName,
		"LastName", u.LastName,
		"UserName", u.UserName,

		"SessionId", u.SessionId,
		"InternalId", u.InternalId,
		"UserGroup", u.UserGroup,
		"Teir", u.Teir,

		"Email", u.Email,
		"Phone", u.Phone,
		"Address", u.Address,
		"Location", u.Location,
		"Url", u.Url,

		"Created_At", u.Created_At,
		"Updated_At", u.Updated_At,
		"Updated_By", u.Updated_By)

	if err != nil {
		fmt.Println(err)
	}

	return u, n, err
}
