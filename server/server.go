package main

import (
	"encoding/json"
	"fake_data"
	"fmt"
	"http_handling"
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
	http_handling.ServeStaticFiles(
		"/views/",
		"/assets/")

	// Handle serving the page -- mapping 'urls' to 'views'
	http.HandleFunc("/app/get-users", fake_data.GetUsers)

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
