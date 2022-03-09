package main

import (
	v4 "github.com/swaggest/swgui/v4emb"
	"net/http"
	// "github.com/swaggest/swgui/v3" // For go1.15 and below.
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/docs/", v4.NewHandler("My API", "/openapi.json", "/docs"))
	http.Handle("/docs/", mux)

	http.ListenAndServe(":8080", nil)
}

//
//func main() {
//	admin := http.NewServeMux()
//
//	// Prefix all paths with the mount point. A ServeMux matches
//	// the full path, even when invoked from another ServeMux.
//	mountPoint := "/admin"
//	admin.HandleFunc(mountPoint+"/", root)
//	admin.HandleFunc(mountPoint+"/foo", foo)
//
//	// Add a trailing "/" to the mount point to indicate a subtree match.
//	http.Handle(mountPoint+"/", admin)
//
//	http.HandleFunc("/ping", pong)
//
//	http.ListenAndServe(":4567", nil)
//}
//
