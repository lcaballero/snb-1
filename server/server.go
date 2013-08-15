package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	// Locations where static files reside.
	ServeStaticFiles("/public/", "/views/")

	http.HandleFunc("/api/all-users/", getAllUser)

	// Start Server on Port
	port := strconv.FormatInt(8080, 10)
	fmt.Println("Starting server on ", port)
	http.ListenAndServe(":"+port, nil)
}

