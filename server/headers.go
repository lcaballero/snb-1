// package main

// import (
// 	"net/http"
// 	"fmt"
// 	"io/ioutil"
// 	"path"
// 	"strings"
// )

// func ServeStaticFiles(roots... string) {
// 	for _, root := range roots {
// 		http.HandleFunc(root, serveFiles(root))
// 	}
// }

// func serveFiles(roots... string) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		file := r.URL.Path

// 		if file[0:1] == "/" {
// 			file = file[1:]
// 		}

// 		fmt.Println("Path: ", r.URL.Path, " Serving File: ", file)

// 		found := false;

// 		for _, root := range roots {

// 			if strings.HasPrefix(r.URL.Path, root) {
// 				toContentType(file, w)
// 				writeFileResponse(file, w)
// 				found = true;
// 				break;
// 			}
// 		}

// 		if !found {
// 			fmt.Println("Couldn't find file: ", file)
// 		}
// 	}
// }

// func getContentType(file string) (string, bool) {
// 	ext := path.Ext(file)

// 	switch ext {
// 	case ".js":
// 		return "text/javascript", false
// 	case ".css":
// 		return "text/css", false
// 	case ".html":
// 		return "text/html", false
// 	case ".json":
// 		return "application/json", false
// 	case ".png":
// 		return "image/png", true
// 	case ".jpeg", ".jpg":
// 		return "image/jpeg", true
// 	}

// 	return "", false;
// }

// func toContentType(file string, w http.ResponseWriter) {
// 	contentType, isImage := getContentType(file)
// 	h := w.Header()

// 	if isImage {
// 		h["Content-Type"] = []string { contentType }	
// 	} else {
// 		h["Content-Type"] = []string { contentType, "charset-utf-8", }	
// 	}
// }

// func writeFileResponse(file string, w http.ResponseWriter) {

// 	bytes, err := ioutil.ReadFile(file)

// 	if err == nil {
// 		fmt.Fprint(w, string(bytes))	
// 	} else {
// 		fmt.Fprint(w, err)
// 	}	
// }

// func serveFile(file string) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		toContentType(file, w)
// 		writeFileResponse(file, w)
// 	}
// }
