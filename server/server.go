package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	// Locations where static files reside.
	ServeStaticFiles("/public/", "/views/")

	http.HandleFunc("/api/getUser/", getUser)


	// Start Server on Port
	port := strconv.FormatInt(8080, 10)
	fmt.Println("Starting server on ", port)
	http.ListenAndServe(":"+port, nil)
}

// ======================= Test ======================= //

func getUser(w http.ResponseWriter, r *http.Request){
	h := w.Header()
	h["Content-Type"] = []string { "application/json", "charset-utf-8", }	

	s := "[{name:'John', email:'John Smith'}]"

	fmt.Fprint(w, s)
}
