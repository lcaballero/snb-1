package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	// Locations where static files reside.
	ServeStaticFiles(
		"/css/",
		"/app/",
		"/lib/angular/")

	// Handle serving the page -- mapping 'urls' to 'views'
	http.HandleFunc("/app/index.html", serveFile("views/index.html"))

	// Start Server on Port
	port := strconv.FormatInt(8080, 10)
	fmt.Println("Starting server on ", port)
	http.ListenAndServe(":"+port, nil)
}
