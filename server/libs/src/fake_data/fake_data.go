package fake_data

import (
	"strconv"
)

type User struct {
	Name  string
	Email string
	Id    string
}

func GetUsers() []User {
	a := make([]User, 10)

	for i := 0; i < 10; i++ {
		n := strconv.Itoa(i)
		u := User{"name-" + n, "email-" + n, n}
		a[i] = u
	}

	return a
}

func FindUser(accept func(u User) bool) (User, bool) {
	var u User
	users := GetUsers()
	for i := 0; i < len(users); i++ {
		u = users[i]
		if accept(u) {
			return u, true
		}
	}
	return u, false
}
