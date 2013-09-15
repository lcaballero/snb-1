package fake_data

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	h := w.Header()
	h["Content-Type"] = []string{"application/json"}

	a := ReadUsers()

	j, _ := json.Marshal(a)
	s := string(j)

	fmt.Fprint(w, s)
}

type Message struct {
	Success bool
	Data    interface{}
}

func FindUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	h := w.Header()
	h["Content-Type"] = []string{"application/json"}

	name := r.Form.Get("name")
	email := r.Form.Get("email")

	var accept func(User) bool

	if len(name) > 0 {
		accept = ByUsername(name)
	} else if len(email) > 0 {
		accept = ByEmail(email)
	} else {
		// return the first user
		accept = func(User) bool { return true }
	}

	u, found := WhereUser(accept)

	var msg Message

	if found {
		msg = Message{found, u}
	} else {
		msg = Message{Success: found}
	}

	j, _ := json.Marshal(msg)
	s := string(j)
	fmt.Fprint(w, s)

}

func ByUsername(name string) func(User) bool {
	return func(u User) bool {
		return u.Name == name
	}
}

func ByEmail(email string) func(User) bool {
	return func(u User) bool {
		return u.Email == email
	}
}
