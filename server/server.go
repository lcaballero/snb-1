package main

import (
	"encoding/json"
	"fake_data"
	"fmt"
	"net/http"
	"strconv"
)

type Response map[string]interface{}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

func main() {

	// Locations where static files reside.
	ServeStaticFiles(
		"/views/",
		"/assets/")

	// Handle serving the page -- mapping 'urls' to 'views'
	http.HandleFunc("/app/get-users", GetUsers)

	// Start Server on Port
	port := strconv.FormatInt(8080, 10)
	fmt.Println("Starting server on ", port)
	http.ListenAndServe(":"+port, nil)
}

// ======================= Test ======================= //

func getUser(w http.ResponseWriter, r *http.Request) {
	//h := w.Header()
	//h["Content-Type"] = []string { "application/json", "charset-utf-8", }
	w.Header().Set("Content-Type", "application/json")
	//s := "[{name:'John', email:'John Smith'}]"

	fmt.Fprint(w, Response{"success": true, "msg": "[{name:'John', email:'John Smith'}]"})
}

func addUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	email := r.FormValue("email")
	pw := r.FormValue("pw")

	fmt.Println(email + ": " + pw)

	CreatUser(email, pw)

	user := ReadUserFromEmail(email)
	fmt.Fprint(w, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	h := w.Header()
	h["Content-Type"] = []string{"application/json"}

	a := fake_data.GetUsers()

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

	var accept func(fake_data.User) bool

	if len(name) > 0 {
		accept = ByUsername(name)
	} else if len(email) > 0 {
		accept = ByEmail(email)
	} else {
		// return the first user
		accept = func(fake_data.User) bool { return true }
	}

	u, found := fake_data.FindUser(accept)

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

func ByUsername(name string) func(fake_data.User) bool {
	return func(u fake_data.User) bool {
		return u.Name == name
	}
}

func ByEmail(email string) func(fake_data.User) bool {
	return func(u fake_data.User) bool {
		return u.Email == email
	}
}
